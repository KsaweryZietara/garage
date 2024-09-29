import {ActivityIndicator, FlatList, StatusBar, Text, View} from "react-native";
import React, {useEffect, useState} from "react";
import axios from "axios";
import {get} from "@/utils/auth";
import EmailDisplay from "@/components/EmailDisplay";
import MenuModal from "@/components/MenuModal";
import {getEmail} from "@/utils/jwt";
import TabSwitcher from "@/components/TabSwitcher";

interface Appointments {
    upcoming: Appointment[];
    inProgress: Appointment[];
    completed: Appointment[];
}

interface Appointment {
    id: number;
    startTime: Date;
    endTime: Date;
    service: Service;
    employee: Employee;
    garage: Garage;
}

interface Garage {
    id: number;
    name: string;
    city: string;
    street: string;
    number: string;
    postalCode: string;
    phoneNumber: string;
}

interface Employee {
    id: number;
    name: string;
    surname: string;
}

interface Service {
    id: number;
    name: string;
    time: string;
    price: string;
}

const AppointmentsScreen = () => {
    const [email, setEmail] = useState<string | null>(null);
    const [menuVisible, setMenuVisible] = useState(false);
    const [appointments, setAppointments] = useState<Appointments | null>(null);
    const [loadingAppointments, setLoadingAppointments] = useState<boolean>(true);
    const [activeTab, setActiveTab] = useState<"upcoming" | "inProgress" | "completed">("upcoming");

    useEffect(() => {
        const fetchEmail = async () => {
            const email = await getEmail("customer_jwt");
            setEmail(email);
        };

        const fetchAppointments = async () => {
            setLoadingAppointments(true);
            try {
                const token = await get("customer_jwt");
                const response = await axios.get<Appointments>("/api/customer/appointments", {
                    headers: {"Authorization": `Bearer ${token}`}
                });
                setAppointments(response.data);
            } catch (error) {
                console.error(error);
            } finally {
                setLoadingAppointments(false);
            }
        };

        fetchAppointments();
        fetchEmail();
    }, []);

    const formatDateTime = (date: Date): string => {
        const d = new Date(date);
        const formattedDate = `${d.getUTCDate().toString().padStart(2, '0')}/${(d.getUTCMonth() + 1).toString().padStart(2, '0')}/${d.getUTCFullYear()}`;
        const hours = d.getUTCHours().toString().padStart(2, '0');
        const minutes = d.getUTCMinutes().toString().padStart(2, '0');
        const formattedTime = `${hours}:${minutes}`;
        return `${formattedTime} ${formattedDate}`;
    };

    const renderAppointmentItem = ({item}: { item: Appointment }) => {
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
