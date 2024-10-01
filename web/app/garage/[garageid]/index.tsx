import {useLocalSearchParams, useRouter} from "expo-router";
import React, {useEffect, useState} from "react";
import {ActivityIndicator, FlatList, StatusBar, Text, TouchableOpacity, View} from "react-native";
import axios from "axios";
import {getEmail} from "@/utils/jwt";
import EmailDisplay from "@/components/EmailDisplay";
import MenuModal from "@/components/MenuModal";
import {Garage, Service} from "@/types";
import {CUSTOMER_JWT} from "@/constants/constants";

const GarageScreen = () => {
    const router = useRouter();
    const [email, setEmail] = useState<string | null>(null);
    const [menuVisible, setMenuVisible] = useState(false);
    const {garageid} = useLocalSearchParams();
    const [garage, setGarage] = useState<Garage | null>(null);
    const [services, setServices] = useState<Service[]>([]);
    const [loadingGarage, setLoadingGarage] = useState<boolean>(true);
    const [loadingServices, setLoadingServices] = useState<boolean>(true);

    useEffect(() => {
        const fetchEmail = async () => {
            const email = await getEmail(CUSTOMER_JWT);
            setEmail(email);
        };

        fetchEmail();
    }, []);

    useEffect(() => {
        const fetchGarage = async () => {
            setLoadingGarage(true);
            await axios.get<Garage>(`/api/garages/${garageid}`)
                .then((response) => {
                    if (response.data) {
                        setGarage(response.data);
                    }
                })
                .catch((error) => {
                    console.error(error);
                })
                .finally(() => {
                    setLoadingGarage(false);
                });
        };

        const fetchServices = async () => {
            setLoadingServices(true);
            await axios.get<Service[]>(`/api/garages/${garageid}/services`)
                .then((response) => {
                    if (response.data) {
                        setServices(response.data);
                    }
                })
                .catch((error) => {
                    console.error(error);
                }).finally(() => {
                    setLoadingServices(false);
                });
        };

        fetchGarage();
        fetchServices();
    }, [garageid]);

    const renderServiceItem = ({item}: { item: Service }) => (
        <TouchableOpacity
            onPress={() => router.push(`/garage/${garageid}/service/${item.id}`)}
            className="p-4 my-2 mx-4 bg-[#2d2d2d] rounded-lg border border-[#444]"
        >
            <View className="p-4 my-2 mx-4 bg-[#2d2d2d] rounded-lg">
                <Text className="text-lg font-bold text-white">{item.name}</Text>
                <Text className="text-[#ddd]">Czas: {item.time}</Text>
                <Text className="text-[#bbb]">Cena: {item.price}</Text>
            </View>
        </TouchableOpacity>
    );

    return (
        <View className="flex-1 bg-black">
            <View className="flex-row justify-between p-4 bg-black">
                <Text className="text-white text-2xl font-bold">GARAGE</Text>
                <EmailDisplay email={email} setMenuVisible={setMenuVisible}/>
            </View>

            {loadingGarage ? (
                <ActivityIndicator size="large" color="#ff5c5c"/>
            ) : (
                garage && (
                    <View className="p-6 bg-[#1a1a1a] rounded-lg mx-4 mt-4 shadow-lg">
                        <Text className="text-3xl font-extrabold text-white mb-2">{garage.name}</Text>
                        <Text className="text-xl text-[#ddd] mb-1">{garage.street} {garage.number}</Text>
                        <Text className="text-lg text-[#aaa] mb-1">{garage.city}, {garage.postalCode}</Text>
                        <Text className="text-lg text-[#aaa]">Telefon: {garage.phoneNumber}</Text>
                    </View>
                )
            )}

            <Text className="text-white text-2xl font-bold mt-8 ml-4 mb-3">Usługi</Text>

            {loadingServices ? (
                <ActivityIndicator size="large" color="#ff5c5c"/>
            ) : (
                <FlatList
                    data={services}
                    renderItem={renderServiceItem}
                    keyExtractor={(item, index) => item?.id?.toString() || index.toString()}
                    ListEmptyComponent={
                        <View className="flex-1 justify-center items-center mb-40">
                            <Text className="text-white text-xl">Brak dostępnych usług.</Text>
                        </View>
                    }
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

export default GarageScreen;
