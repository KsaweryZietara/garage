import {useLocalSearchParams, useRouter} from "expo-router";
import React, {useEffect, useState} from "react";
import {ActivityIndicator, FlatList, StatusBar, Text, View} from "react-native";
import axios from "axios";

interface Service {
    id: number;
    name: string;
    time: string;
    price: string;
}

interface Employee {
    id: number;
    name: string;
    surname: string;
}

const ServiceScreen = () => {
    const router = useRouter();
    const {garageid, serviceid} = useLocalSearchParams();
    const [service, setService] = useState<Service | null>(null);
    const [employees, setEmployees] = useState<Employee[]>([]);
    const [loadingService, setLoadingService] = useState<boolean>(true);
    const [loadingEmployees, setLoadingEmployees] = useState<boolean>(true);

    useEffect(() => {
        const fetchService = async () => {
            setLoadingService(true);
            try {
                const response = await axios.get<Service>(`/api/services/${serviceid}`);
                if (response.data) {
                    setService(response.data);
                }
            } catch (error) {
                console.error(error);
            } finally {
                setLoadingService(false);
            }
        };

        const fetchEmployees = async () => {
            setLoadingEmployees(true);
            try {
                const response = await axios.get<Employee[]>(`/api/garages/${garageid}/employees`);
                if (response.data) {
                    setEmployees(response.data);
                }
            } catch (error) {
                console.error(error);
            } finally {
                setLoadingEmployees(false);
            }
        };

        fetchService();
        fetchEmployees();
    }, [garageid, serviceid]);

    const renderEmployeeItem = ({item}: { item: Employee }) => (
        <View className="p-4 my-2 mx-4 bg-[#2d2d2d] rounded-lg border border-[#444]">
            <Text className="text-lg font-bold text-white">{item.name} {item.surname}</Text>
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

            {loadingService ? (
                <ActivityIndicator size="large" color="#ff5c5c"/>
            ) : (
                service && (
                    <View className="p-6 bg-[#1a1a1a] rounded-lg mx-4 mt-4 shadow-lg">
                        <Text className="text-3xl font-extrabold text-white mb-2">{service.name}</Text>
                        <Text className="text-xl text-[#ddd] mb-1">Czas trwania: {service.time}</Text>
                        <Text className="text-lg text-[#bbb]">Cena: {service.price}</Text>
                    </View>
                )
            )}

            <Text className="text-white text-2xl font-bold mt-8 ml-4 mb-3">Pracownicy</Text>

            {loadingEmployees ? (
                <ActivityIndicator size="large" color="#ff5c5c"/>
            ) : (
                <FlatList
                    data={employees}
                    renderItem={renderEmployeeItem}
                    keyExtractor={(item) => item.id.toString()}
                    ListEmptyComponent={
                        <View className="flex-1 justify-center items-center mb-40">
                            <Text className="text-white text-xl">Brak dostępnych pracowników.</Text>
                        </View>
                    }
                    showsHorizontalScrollIndicator={false}
                />
            )}

            <StatusBar backgroundColor="#000000"/>
        </View>
    );
};

export default ServiceScreen;
