import {useGlobalSearchParams, useRouter} from "expo-router";
import React, {useEffect, useState} from "react";
import {ActivityIndicator, FlatList, StatusBar, Text, TouchableOpacity, View} from "react-native";
import {Searchbar} from "react-native-paper";
import axios from "axios";

interface Garage {
    id: number;
    name: string;
    city: string;
    street: string;
    number: string;
    postalCode: string;
    phoneNumber: string;
}

const GaragesScreen = () => {
    const router = useRouter();
    const query = useGlobalSearchParams();
    const [search, setSearch] = useState<string>("");
    const [garages, setGarages] = useState<Garage[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [page, setPage] = useState<number>(1);
    const [hasMore, setHasMore] = useState<boolean>(true);
    const [isRefreshing, setIsRefreshing] = useState<boolean>(false);

    const handleSearch = () => {
        router.push({pathname: "/garages", params: {search: search}});
    };

    const fetchGarages = (query: string | string[], page = 1) => {
        if (!hasMore && page !== 1) return;

        setLoading(true);

        axios.get<Garage[]>("/api/garages", {
            params: {
                query,
                page,
            },
        })
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
    }, [page]);

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
            <Text className="text-[#ddd]">{item.street} {item.number}</Text>
            <Text className="text-[#bbb]">{item.city}, {item.postalCode}</Text>
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
                <Text className="mt-2 text-[#ff5c5c] font-bold" onPress={() => router.push("/login")}>
                    ZALOGUJ SIĘ
                </Text>
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

            <StatusBar backgroundColor="#000000"/>
        </View>
    );
};

export default GaragesScreen;
