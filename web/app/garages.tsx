import {useGlobalSearchParams, useRouter} from "expo-router";
import React, {useEffect, useState} from "react";
import {ActivityIndicator, FlatList, StatusBar, Text, TouchableOpacity, View} from "react-native";
import {Searchbar} from "react-native-paper";
import axios from "axios";
import {getEmail} from "@/utils/jwt";
import EmailDisplay from "@/components/EmailDisplay";
import MenuModal from "@/components/MenuModal";
import {Garage} from "@/types";
import {CUSTOMER_JWT} from "@/constants/constants";
import {AirbnbRating} from "react-native-ratings";
import * as Location from 'expo-location';
import {LocationObject} from "expo-location";

const GaragesScreen = () => {
    const router = useRouter();
    const [email, setEmail] = useState<string | null>(null);
    const [menuVisible, setMenuVisible] = useState(false);
    const query = useGlobalSearchParams();
    const [search, setSearch] = useState<string>("");
    const [garages, setGarages] = useState<Garage[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [page, setPage] = useState<number>(1);
    const [hasMore, setHasMore] = useState<boolean>(true);
    const [isRefreshing, setIsRefreshing] = useState<boolean>(false);
    const [location, setLocation] = useState<LocationObject | null>(null);

    useEffect(() => {
        const fetchEmail = async () => {
            const email = await getEmail(CUSTOMER_JWT);
            setEmail(email);
        };

        fetchEmail();
    }, []);

    useEffect(() => {
        const getLocation = async () => {
            try {
                const {status} = await Location.requestForegroundPermissionsAsync();
                if (status !== 'granted') {
                    console.warn('Permission to access location was denied');
                    return;
                }

                const currentLocation = await Location.getCurrentPositionAsync({});
                setLocation(currentLocation);
            } catch (error) {
                console.error("Error getting location:", error);
            }
        };

        getLocation();
    }, []);

    const handleSearch = () => {
        router.push({pathname: "/garages", params: {search: search}});
    };

    const fetchGarages = (query: string | string[], page = 1) => {
        if (!hasMore && page !== 1) return;

        setLoading(true);

        const params: any = {
            query,
            page,
        };

        if (location) {
            params.latitude = location.coords.latitude;
            params.longitude = location.coords.longitude;
        }

        axios.get<Garage[]>("/api/garages", {params})
            .then((response) => {
                const fetchedGarages = response.data;

                if (page === 1) {
                    setGarages(fetchedGarages);
                } else {
                    setGarages(prevGarages => [...prevGarages, ...fetchedGarages]);
                }

                if (fetchedGarages.length === 0) {
                    setHasMore(false);
                }
            })
            .catch((error) => {
                console.error(error);
            })
            .finally(() => {
                setLoading(false);
                setIsRefreshing(false);
            });
    };

    useEffect(() => {
        fetchGarages(query.search, page);
    }, [page, location]);

    const loadMoreGarages = () => {
        if (!loading && hasMore) {
            setPage(prevPage => prevPage + 1);
        }
    };

    const onRefresh = () => {
        setIsRefreshing(true);
        setPage(1);
        setHasMore(true);
        fetchGarages(query.search, 1);
    };

    const renderGarageItem = ({item}: { item: Garage }) => (
        <TouchableOpacity
            onPress={() => router.push(`/garage/${item.id}`)}
            className="p-4 my-2 mx-4 bg-[#2d2d2d] rounded-lg border border-[#444]"
        >
            <Text className="text-xl font-bold text-white">{item.name}</Text>
            {item.rating !== 0 && (
                <View style={{flexDirection: 'row', alignItems: 'center'}}>
                    <Text className="text-[#ddd] text-lg">{item.rating}</Text>
                    <AirbnbRating
                        isDisabled={true}
                        count={5}
                        defaultRating={item.rating}
                        size={15}
                        showRating={false}
                        selectedColor="#ef4444"
                        starContainerStyle={{marginLeft: 3}}
                    />
                </View>
            )}
            <Text className="text-[#ddd]">{item.street} {item.number}</Text>
            <Text className="text-[#ddd]">{item.city}, {item.postalCode}</Text>
            {item.distance !== 0 && (
                <Text className="text-[#ddd] mt-1">{item.distance} km</Text>
            )}
        </TouchableOpacity>
    );

    const renderEmptyMessage = () => (
        <View className="flex-1 justify-center items-center mb-40">
            <Text className="text-white text-xl">Przykro nam, niczego nie znaleźliśmy.</Text>
        </View>
    );

    return (
        <View className="flex-1 bg-black">
            <View className="flex-row justify-between p-4 bg-black">
                <Text className="text-white text-2xl font-bold">GARAGE</Text>
                <EmailDisplay email={email} setMenuVisible={setMenuVisible}/>
            </View>

            <View className="items-center mb-4">
                <Searchbar
                    className="w-4/5 lg:w-2/5 bg-[#2d2d2d] col-white"
                    placeholder="Szukaj warsztatów lub usługi"
                    placeholderTextColor="#aaa"
                    inputStyle={{color: "white"}}
                    onChangeText={setSearch}
                    onIconPress={handleSearch}
                    onSubmitEditing={handleSearch}
                    value={search}
                />
            </View>

            {garages.length === 0 && !loading ? (
                renderEmptyMessage()
            ) : (
                <FlatList
                    data={garages}
                    renderItem={renderGarageItem}
                    keyExtractor={(item) => item.id.toString()}
                    onEndReached={loadMoreGarages}
                    onEndReachedThreshold={0.5}
                    ListFooterComponent={loading ? <ActivityIndicator size="large" color="#ff5c5c"/> : null}
                    refreshing={isRefreshing}
                    onRefresh={onRefresh}
                    showsHorizontalScrollIndicator={false}
                />
            )}

            <MenuModal
                visible={menuVisible}
                onClose={() => setMenuVisible(false)}
                email={email}
                setEmail={setEmail}
            />

            <StatusBar backgroundColor="#000000"/>
        </View>
    );
};

export default GaragesScreen;
