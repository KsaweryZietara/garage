import {ActivityIndicator, FlatList, StatusBar, Text, View} from "react-native";
import React, {useEffect, useState} from "react";
import axios from "axios";
import {get} from "@/utils/auth";
import EmailDisplay from "@/components/EmailDisplay";
import MenuModal from "@/components/MenuModal";
import {getEmail} from "@/utils/jwt";
import TabSwitcher from "@/components/TabSwitcher";
import {CustomerAppointment, Appointments} from "@/types";
import {CUSTOMER_JWT} from "@/constants/constants";
import {formatDateTime} from "@/utils/time";

const AppointmentsScreen = () => {
    const [email, setEmail] = useState<string | null>(null);
    const [menuVisible, setMenuVisible] = useState(false);
    const [appointments, setAppointments] = useState<Appointments | null>(null);
    const [loadingAppointments, setLoadingAppointments] = useState<boolean>(true);
    const [activeTab, setActiveTab] = useState<"upcoming" | "inProgress" | "completed">("upcoming");

    useEffect(() => {
        const fetchEmail = async () => {
            const email = await getEmail(CUSTOMER_JWT);
            setEmail(email);
        };

        const fetchAppointments = async () => {
            setLoadingAppointments(true);
            const token = await get(CUSTOMER_JWT);
            await axios.get<Appointments>("/api/customers/appointments", {
                headers: {"Authorization": `Bearer ${token}`}
            })
                .then((response) => {
                    setAppointments(response.data);
                })
                .catch((error) => {
                    console.error(error);
                })
                .finally(() => {
                    setLoadingAppointments(false);
                });
        };

        fetchAppointments();
        fetchEmail();
    }, []);

    const renderAppointmentItem = ({item}: { item: CustomerAppointment }) => {
        return (
            <View className="p-4 my-2 mx-4 bg-[#2d2d2d] rounded-lg">
                <Text className="text-lg text-white font-bold">
                    {item.service.name}
                </Text>
                <Text className="text-sm text-[#ddd]">
                    {formatDateTime(item.startTime)} - {formatDateTime(item.endTime)}
                </Text>
                <Text className="text-sm text-[#ddd]">
                    Cena: {item.service.price}
                </Text>
                <Text className="text-sm text-[#ddd]">
                    {item.garage.name}
                </Text>
                <Text className="text-sm text-[#ddd]">
                    ul. {item.garage.street} {item.garage.number}
                </Text>
                <Text className="text-sm text-[#ddd]">
                    {item.garage.postalCode} {item.garage.city}
                </Text>
                <Text className="text-sm text-[#ddd]">
                    {item.garage.phoneNumber}
                </Text>
                <Text className="text-sm text-[#ddd]">
                    Mechanik: {item.employee.name} {item.employee.surname}
                </Text>
            </View>
        );
    };

    const getCurrentAppointments = () => {
        switch (activeTab) {
            case "upcoming":
                return appointments?.upcoming ?? [];
            case "inProgress":
                return appointments?.inProgress ?? [];
            case "completed":
                return appointments?.completed ?? [];
            default:
                return [];
        }
    };

    return (
        <View className="flex-1 bg-black">
            <View className="flex-row justify-between p-4 bg-black">
                <Text className="text-white text-2xl lg:text-4xl font-bold">GARAGE</Text>
                <EmailDisplay email={email} setMenuVisible={setMenuVisible}/>
            </View>

            <TabSwitcher activeTab={activeTab} setActiveTab={setActiveTab}/>

            {loadingAppointments ? (
                <ActivityIndicator size="large" color="#ff5c5c"/>
            ) : (
                <FlatList
                    data={getCurrentAppointments()}
                    renderItem={renderAppointmentItem}
                    keyExtractor={(item, index) => index.toString()}
                    ListEmptyComponent={
                        <View className="flex-1 justify-end items-center mb-40">
                            <Text className="text-white text-xl">Brak wizyt w tej sekcji.</Text>
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

export default AppointmentsScreen;
