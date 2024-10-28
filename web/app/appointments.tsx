import {
    ActivityIndicator,
    FlatList,
    StatusBar,
    Text,
    View,
    Modal,
    TextInput,
    TouchableWithoutFeedback,
} from "react-native";
import React, {useEffect, useState} from "react";
import axios from "axios";
import {get} from "@/utils/auth";
import EmailDisplay from "@/components/EmailDisplay";
import MenuModal from "@/components/MenuModal";
import {getJwtPayload} from "@/utils/jwt";
import TabSwitcher from "@/components/TabSwitcher";
import {CustomerAppointment, Appointments} from "@/types";
import {CUSTOMER_JWT} from "@/constants/constants";
import {formatDateTime} from "@/utils/time";
import CustomButton from "@/components/CustomButton";
import {AirbnbRating} from "react-native-ratings";
import {useRouter} from "expo-router";

const AppointmentsScreen = () => {
    const router = useRouter();
    const [email, setEmail] = useState<string | null>(null);
    const [menuVisible, setMenuVisible] = useState(false);
    const [appointments, setAppointments] = useState<Appointments | null>(null);
    const [loadingAppointments, setLoadingAppointments] = useState<boolean>(true);
    const [activeTab, setActiveTab] = useState<"upcoming" | "inProgress" | "completed">("upcoming");

    const [reviewVisible, setReviewVisible] = useState<boolean>(false);
    const [comment, setComment] = useState<string>("");
    const [rating, setRating] = useState<number>(3);
    const [selectedAppointmentId, setSelectedAppointmentId] = useState<number | null>(null);
    const [reviewSubmitted, setReviewSubmitted] = useState<boolean>(false);

    useEffect(() => {
        fetchAppointments();
        fetchEmail();
    }, []);

    const fetchEmail = async () => {
        const jwtPayload = await getJwtPayload(CUSTOMER_JWT);
        setEmail(jwtPayload?.email || null);
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

    const handleSubmit = async () => {
        if (!selectedAppointmentId) return;

        const token = await get(CUSTOMER_JWT);
        await axios.put(`/api/appointments/${selectedAppointmentId}/reviews`, {rating, comment}, {
            headers: {"Authorization": `Bearer ${token}`}
        })
            .then(() => {
                setReviewSubmitted(true);
                fetchAppointments();
            })
            .catch((error) => {
                console.error(error);
            }).finally(() => {
                setRating(3)
                setComment("")
            });
    };

    const handleDelete = async () => {
        if (!selectedAppointmentId) return;

        const token = await get(CUSTOMER_JWT);
        await axios.delete(`/api/appointments/${selectedAppointmentId}/reviews`, {
            headers: {"Authorization": `Bearer ${token}`}
        })
            .then(() => {
                setReviewVisible(false);
                fetchAppointments();
            })
            .catch((error) => {
                console.error(error);
            }).finally(() => {
                setRating(3)
                setComment("")
            });
    };

    const handleAppointmentDelete = async (id: number) => {
        const token = await get(CUSTOMER_JWT);
        await axios.delete(`/api/appointments/${id}`, {
            headers: {"Authorization": `Bearer ${token}`}
        })
            .then(() => {
                fetchAppointments();
            })
            .catch((error) => {
                console.error(error);
            });
    };

    const renderAppointmentItem = ({item}: { item: CustomerAppointment }) => {
        const timeUntilStart = new Date(item.startTime).getTime() - new Date().getTime();
        const isMoreThan24Hours = timeUntilStart > 24 * 60 * 60 * 1000;

        return (
            <View className="p-4 my-2 mx-4 bg-[#2d2d2d] rounded-lg">
                <View className="flex-col lg:flex-row justify-between lg:items-center">
                    <View className="flex-1">
                        <Text className="text-xl text-white font-bold">
                            {item.service.name}
                        </Text>
                        <Text className="text-sm text-[#ddd]">
                            {formatDateTime(item.startTime)} - {formatDateTime(item.endTime)}
                        </Text>
                        <Text className="text-sm text-[#ddd]">
                            Cena: {item.service.price} zł
                        </Text>
                        <Text className="text-sm text-[#ddd]">
                            Samochód: {item.car.make} {item.car.model}
                        </Text>
                        <Text className="text-sm text-[#ddd]">
                            Mechanik: {item.employee.name} {item.employee.surname}
                        </Text>
                        <Text className="text-sm text-[#ddd]">
                            Telefon: {item.garage.phoneNumber}
                        </Text>
                        <Text className="text-lg text-[#ddd] mt-1">
                            Adres:
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
                    </View>

                    {activeTab === "completed" && (
                        <CustomButton
                            title={item.rating ? "Edytuj opinie" : "Dodaj opinie"}
                            onPress={() => {
                                setRating(item.rating ?? 3);
                                setComment(item.comment ?? "");
                                setSelectedAppointmentId(item.id);
                                setReviewSubmitted(false);
                                setReviewVisible(true);
                            }}
                            containerStyles="bg-red-500 self-center mt-3 lg:mt-0 w-2/5 lg:w-1/5"
                            textStyles="text-white font-bold"
                        />
                    )}
                    {activeTab === "upcoming" && isMoreThan24Hours && (
                        <CustomButton
                            title={"Anuluj wizytę"}
                            onPress={() => {
                                handleAppointmentDelete(item.id)
                            }}
                            containerStyles="bg-red-500 self-center mt-3 lg:mt-0 w-2/5 lg:w-1/5"
                            textStyles="text-white font-bold"
                        />
                    )}
                </View>
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
                <Text className="text-white text-2xl lg:text-4xl font-bold lg:mt-1.5"
                      onPress={() => router.push("/home")}>GARAGE</Text>
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

            <Modal
                visible={reviewVisible}
                animationType="fade"
                transparent={true}
                onRequestClose={() => setReviewVisible(false)}
            >
                <TouchableWithoutFeedback onPress={() => setReviewVisible(false)}>
                    <View className="flex-1 justify-center items-center"
                          style={{backgroundColor: 'rgba(0, 0, 0, 0.75)'}}>
                        <TouchableWithoutFeedback onPress={() => {
                        }}>
                            <View className="bg-[#2d2d2d] p-5 rounded-lg w-4/5 lg:w-2/5">
                                {reviewSubmitted ? (
                                    <Text className="text-white self-center text-lg font-bold my-5">
                                        Dziękujemy za twoją opinie!
                                    </Text>) : (
                                    <View>
                                        <Text className="text-white text-lg font-bold mb-5">Dodaj swoją opinię</Text>
                                        <AirbnbRating
                                            count={5}
                                            defaultRating={rating}
                                            size={40}
                                            showRating={false}
                                            selectedColor="#ef4444"
                                            onFinishRating={setRating}
                                        />
                                        <TextInput
                                            value={comment}
                                            onChangeText={setComment}
                                            placeholder="Wpisz swoją opinię"
                                            multiline
                                            numberOfLines={4}
                                            className="border p-2 mt-5 rounded text-#2d2d2d bg-white align-text-top max-h-20"
                                            placeholderTextColor="#2d2d2d"
                                        />
                                        <CustomButton
                                            title="Wyślij"
                                            onPress={handleSubmit}
                                            containerStyles="bg-red-500 self-center mt-5 px-10 lg:px-20"
                                            textStyles="text-white font-bold"
                                        />
                                        <Text className="text-red-500 mt-3 self-center" onPress={handleDelete}>
                                            Usuń opinie
                                        </Text>
                                    </View>
                                )}
                            </View>
                        </TouchableWithoutFeedback>
                    </View>
                </TouchableWithoutFeedback>
            </Modal>

            <StatusBar backgroundColor="#000000"/>
        </View>
    );
};

export default AppointmentsScreen;
