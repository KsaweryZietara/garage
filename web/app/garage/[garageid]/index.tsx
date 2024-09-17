import {useLocalSearchParams, useRouter} from "expo-router";
import React, {useEffect, useState} from "react";
import {ActivityIndicator, FlatList, StatusBar, Text, TouchableOpacity, View} from "react-native";
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

interface Service {
    id: number;
    name: string;
    time: string;
    price: string;
}

const GarageScreen = () => {
    const router = useRouter();
    const {garageid} = useLocalSearchParams();
    const [garage, setGarage] = useState<Garage | null>(null);
    const [services, setServices] = useState<Service[]>([]);
    const [loadingGarage, setLoadingGarage] = useState<boolean>(true);
    const [loadingServices, setLoadingServices] = useState<boolean>(true);

    useEffect(() => {
        const fetchGarage = async () => {
            setLoadingGarage(true);
            try {
                const response = await axios.get<Garage>(`/api/garages/${garageid}`);
                if (response.data) {
                    setGarage(response.data);
                }
            } catch (error) {
                console.error(error);
            } finally {
                setLoadingGarage(false);
            }
        };

        const fetchServices = async () => {
            setLoadingServices(true);
            try {
                const response = await axios.get<Service[]>(`/api/garages/${garageid}/services`);
                if (response.data) {
                    setServices(response.data);
                }
            } catch (error) {
                console.error(error);
            } finally {
                setLoadingServices(false);
            }
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
                <Text className="mt-2 text-[#ff5c5c] font-bold" onPress={() => router.push("/login")}>
                    ZALOGUJ SIĘ
                </Text>
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

            <StatusBar backgroundColor="#000000"/>
        </View>
    );
};

export default GarageScreen;
