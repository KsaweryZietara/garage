import {useLocalSearchParams, useRouter} from "expo-router";
import React, {useEffect, useState} from "react";
import {ActivityIndicator, FlatList, Image, StatusBar, Text, TouchableOpacity, View} from "react-native";
import axios from "axios";
import {getJwtPayload} from "@/utils/jwt";
import EmailDisplay from "@/components/EmailDisplay";
import MenuModal from "@/components/MenuModal";
import {Employee, Service} from "@/types";
import {CUSTOMER_JWT} from "@/constants/constants";

const ServiceScreen = () => {
    const router = useRouter();
    const [email, setEmail] = useState<string | null>(null);
    const [menuVisible, setMenuVisible] = useState(false);
    const {garageid, serviceid} = useLocalSearchParams();
    const [service, setService] = useState<Service | null>(null);
    const [employees, setEmployees] = useState<Employee[]>([]);
    const [loadingService, setLoadingService] = useState<boolean>(true);
    const [loadingEmployees, setLoadingEmployees] = useState<boolean>(true);

    useEffect(() => {
        const fetchEmail = async () => {
            const jwtPayload = await getJwtPayload(CUSTOMER_JWT);
            setEmail(jwtPayload?.email || null);
        };

        fetchEmail();
    }, []);

    useEffect(() => {
        const fetchService = async () => {
            setLoadingService(true);
            await axios.get<Service>(`/api/services/${serviceid}`)
                .then((response) => {
                    if (response.data) {
                        setService(response.data);
                    }
                })
                .catch((error) => {
                    console.error(error);
                })
                .finally(() => {
                    setLoadingService(false);
                });
        };

        const fetchEmployees = async () => {
            setLoadingEmployees(true);
            await axios.get<Employee[]>(`/api/garages/${garageid}/employees`)
                .then((response) => {
                    if (response.data) {
                        setEmployees(response.data);
                    }
                })
                .catch((error) => {
                    console.error(error);
                })
                .finally(() => {
                    setLoadingEmployees(false);
                });
        };

        fetchService();
        fetchEmployees();
    }, [garageid, serviceid]);

    const renderEmployeeItem = ({item}: { item: Employee }) => (
        <TouchableOpacity
            onPress={() => router.push(`/garage/${garageid}/service/${serviceid}/employee/${item.id}`)}
            className="my-2 mx-4 bg-[#2d2d2d] rounded-lg border border-[#444]"
        >
            <View className="p-4 my-2 mx-2 bg-[#2d2d2d] rounded-lg">
                <View style={{flexDirection: 'row', alignItems: 'center'}}>
                    <Image
                        source={{
                            uri: item.profilePicture ?
                                `data:image/png;base64,${item.profilePicture}` :
                                "/assets/profile-picture-placeholder.png"
                        }}
                        style={{width: 60, height: 60, borderRadius: 20, marginRight: 12}}
                    />
                    <View style={{flex: 1}}>
                        <Text className="text-xl font-bold text-white">{item.name} {item.surname}</Text>
                    </View>
                </View>
            </View>
        </TouchableOpacity>
    );

    return (
        <View className="flex-1 bg-black">
            <View className="flex-row justify-between p-4 bg-black">
                <Text className="text-white text-2xl lg:text-4xl font-bold lg:mt-1.5"
                      onPress={() => router.push("/home")}>GARAGE</Text>
                <EmailDisplay email={email} setMenuVisible={setMenuVisible}/>
            </View>

            {loadingService ? (
                <ActivityIndicator size="large" color="#ff5c5c"/>
            ) : (
                service && (
                    <View className="p-6 bg-[#1a1a1a] rounded-lg mx-4 mt-4 shadow-lg">
                        <Text className="text-3xl font-extrabold text-white mb-2">{service.name}</Text>
                        <Text className="text-xl text-[#aaa]">Czas: {service.time} godz.</Text>
                        <Text className="text-xl text-[#aaa]">Cena: {service.price} zł</Text>
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

export default ServiceScreen;
