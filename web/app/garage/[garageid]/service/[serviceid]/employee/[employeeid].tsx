import {useLocalSearchParams, useRouter} from "expo-router";
import React, {useEffect, useState} from "react";
import {ActivityIndicator, FlatList, Modal, Platform, StatusBar, Text, TouchableOpacity, View} from "react-native";
import CalendarStrip from "react-native-calendar-strip";
import axios from "axios";
import moment, {Moment} from "moment";
import "moment/locale/pl";
import {get} from "@/utils/auth";
import {getEmail} from "@/utils/jwt";
import EmailDisplay from "@/components/EmailDisplay";
import MenuModal from "@/components/MenuModal";
import {Employee, Service, TimeSlot} from "@/types";
import {CUSTOMER_JWT} from "@/constants/constants";
import {formatDateTime, formatTime} from "@/utils/time";

moment.locale("pl");

const AppointmentScreen = () => {
    const router = useRouter();
    const [email, setEmail] = useState<string | null>(null);
    const [menuVisible, setMenuVisible] = useState(false);
    const {serviceid, employeeid} = useLocalSearchParams() as { serviceid: string, employeeid: string };
    const [service, setService] = useState<Service | null>(null);
    const [employee, setEmployee] = useState<Employee | null>(null);
    const [timeSlots, setTimeSlots] = useState<TimeSlot[]>([]);
    const [loadingService, setLoadingService] = useState<boolean>(true);
    const [loadingEmployee, setLoadingEmployee] = useState<boolean>(true);
    const [loadingSlots, setLoadingSlots] = useState<boolean>(true);
    const [selectedDate, setSelectedDate] = useState<Moment>(moment());
    const [modalVisible, setModalVisible] = useState<boolean>(false);
    const [selectedTimeSlot, setSelectedTimeSlot] = useState<TimeSlot | null>(null);
    const [feedbackModalVisible, setFeedbackModalVisible] = useState<boolean>(false);
    const [feedbackMessage, setFeedbackMessage] = useState<string>("");

    useEffect(() => {
        const fetchEmail = async () => {
            const email = await getEmail(CUSTOMER_JWT);
            setEmail(email);
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

        const fetchEmployee = async () => {
            setLoadingEmployee(true);
            await axios.get<Employee>(`/api/employees/${employeeid}`)
                .then((response) => {
                    if (response.data) {
                        setEmployee(response.data);
                    }
                })
                .catch((error) => {
                    console.error(error);
                })
                .finally(() => {
                    setLoadingEmployee(false);
                });
        };

        fetchService();
        fetchEmployee();
    }, [serviceid, employeeid]);

    useEffect(() => {
        const fetchAvailableSlots = async () => {
            setLoadingSlots(true);
            await axios.get<TimeSlot[]>(
                `/api/appointments/availableSlots?serviceId=${serviceid}&employeeId=${employeeid}&date=${selectedDate.format("YYYY-MM-DD")}`
            )
                .then((response) => {
                    setTimeSlots(response.data);
                })
                .catch((error) => {
                    console.error(error);
                })
                .finally(() => {
                    setLoadingSlots(false);
                });
        };

        fetchAvailableSlots();
    }, [selectedDate, serviceid, employeeid]);

    const handleSubmit = async () => {
        const token = await get(CUSTOMER_JWT);
        if (token == null) {
            router.push("/login")
            return
        }
        const data = {
            employeeId: parseInt(employeeid, 10),
            serviceId: parseInt(serviceid, 10),
            startTime: selectedTimeSlot?.startTime,
            endTime: selectedTimeSlot?.endTime
        };
        await axios.post("/api/appointments", data, {headers: {"Authorization": `Bearer ${token}`}})
            .then(() => {
                setFeedbackMessage("Rezerwacja zakończona sukcesem!");
            })
            .catch((error) => {
                console.error(error)
                setFeedbackMessage("Wystąpił błąd: " + (error.response?.data?.message || error.message));
            })
            .finally(() => {
                setFeedbackModalVisible(true);
            });
    };

    const handleDateChange = (date: Moment) => {
        setSelectedDate(date);
    };

    const renderTimeSlotItem = ({item}: { item: TimeSlot }) => (
        <TouchableOpacity
            onPress={() => {
                setSelectedTimeSlot(item);
                setModalVisible(true);
            }}
            className="p-4 my-2 mx-4 bg-[#2d2d2d] rounded-lg"
        >
            <Text className="text-lg font-bold text-white">
                {formatTime(item.startTime)}
            </Text>
            <Text className="text-sm text-gray-400">
                {formatDateTime(item.endTime)}
            </Text>
        </TouchableOpacity>
    );

    return (
        <View className="flex-1 bg-black">
            <View className="flex-row justify-between p-4 bg-black">
                <Text className="text-white text-2xl font-bold">GARAGE</Text>
                <EmailDisplay email={email} setMenuVisible={setMenuVisible}/>
            </View>

            {loadingService || loadingEmployee ? (
                <ActivityIndicator size="large" color="#ff5c5c"/>
            ) : (
                service && (
                    <View className="p-6 bg-[#1a1a1a] rounded-lg mx-4 mt-4 shadow-lg">
                        <Text className="text-3xl font-extrabold text-white mb-2">{service.name}</Text>
                        <Text className="text-xl text-[#ddd] mb-1">Cena: {service.price}</Text>
                        <Text className="text-lg text-[#bbb]">Czas trwania: {service.time}</Text>
                        <Text className="text-lg text-[#bbb]">Mechanik: {employee?.name} {employee?.surname}</Text>
                    </View>
                )
            )}

            <CalendarStrip
                scrollable
                locale={{name: "pl", config: {}}}
                calendarAnimation={{type: "sequence", duration: 30}}
                style={{height: 100, paddingTop: 20, paddingBottom: 10}}
                calendarHeaderStyle={{
                    color: "white",
                    marginBottom: Platform.OS === "web" ? 35 : 10,
                    fontSize: 20,
                }}
                calendarColor={"#000"}
                dateNumberStyle={{color: "white"}}
                dateNameStyle={{color: "white"}}
                highlightDateNumberStyle={{color: "#ff5c5c"}}
                highlightDateNameStyle={{color: "#ff5c5c"}}
                disabledDateNameStyle={{color: "gray"}}
                disabledDateNumberStyle={{color: "gray"}}
                iconContainer={{flex: 0.1}}
                markedDates={[
                    {
                        date: moment().format("YYYY-MM-DD"),
                        dots: [{color: "#ff5c5c", selectedColor: "#ff5c5c"}],
                    },
                ]}
                leftSelector={<Text style={{color: "white", fontSize: 30}}>&lt;</Text>}
                rightSelector={<Text style={{color: "white", fontSize: 30}}>&gt;</Text>}
                onDateSelected={handleDateChange}
                selectedDate={selectedDate}
                minDate={moment()}
            />

            <Text className="text-white text-2xl font-bold mt-4 ml-4 mb-3">Dostępne godziny</Text>

            {loadingSlots ? (
                <ActivityIndicator size="large" color="#ff5c5c"/>
            ) : (
                <FlatList
                    data={timeSlots}
                    renderItem={renderTimeSlotItem}
                    keyExtractor={(item, index) => index.toString()}
                    ListEmptyComponent={
                        <View className="flex-1 justify-center items-center mb-40">
                            <Text className="text-white text-xl">Brak dostępnych terminów na dziś.</Text>
                        </View>
                    }
                    showsHorizontalScrollIndicator={false}
                />
            )}

            <Modal
                animationType="fade"
                transparent={true}
                visible={modalVisible}
                onRequestClose={() => setModalVisible(false)}
            >
                <View className="flex-1 justify-center items-center bg-black bg-opacity-50">
                    <View className="w-10/12 lg:w-1/4 p-6 bg-[#1a1a1a] rounded-lg">
                        {selectedTimeSlot && service && employee && (
                            <>
                                <Text className="text-white text-3xl font-bold mb-5 text-center">Podsumowanie</Text>
                                <View className="text-white text-xl mb-2 flex-row justify-between">
                                    <Text className="text-white text-xl mb-2">Usługa:</Text>
                                    <Text className="text-white text-xl mb-2">{service.name}</Text>
                                </View>
                                <View className="text-white text-xl mb-2 flex-row justify-between">
                                    <Text className="text-white text-xl mb-2">Cena:</Text>
                                    <Text className="text-white text-xl mb-2">{service.price}</Text>
                                </View>
                                <View className="text-white text-xl mb-2 flex-row justify-between">
                                    <Text className="text-white text-xl mb-2">Mechanik:</Text>
                                    <Text className="text-white text-xl mb-2">{employee.name} {employee.surname}</Text>
                                </View>
                                <View className="text-white text-xl mb-2 flex-row justify-between">
                                    <Text className="text-white text-xl mb-2">Termin:</Text>
                                    <Text
                                        className="text-white text-xl mb-2">{formatDateTime(selectedTimeSlot.startTime)}</Text>
                                </View>
                                <View className="text-white text-xl mb-4 flex-row justify-between">
                                    <Text className="text-white text-xl mb-2">Odbiór:</Text>
                                    <Text className="text-white text-xl mb-2">
                                        {formatDateTime(selectedTimeSlot.endTime)}
                                    </Text>
                                </View>
                                <TouchableOpacity
                                    className="bg-red-500 p-3 rounded-lg mt-5 mb-4"
                                    onPress={() => {
                                        handleSubmit()
                                        setModalVisible(false);
                                    }}
                                >
                                    <Text className="text-white text-center">Potwierdź rezerwację</Text>
                                </TouchableOpacity>
                                <TouchableOpacity
                                    className="mt-2"
                                    onPress={() => setModalVisible(false)}
                                >
                                    <Text className="text-gray-400 text-center">Anuluj</Text>
                                </TouchableOpacity>
                            </>
                        )}
                    </View>
                </View>
            </Modal>

            <Modal
                animationType="fade"
                transparent={true}
                visible={feedbackModalVisible}
                onRequestClose={() => setFeedbackModalVisible(false)}
            >
                <View className="flex-1 justify-center items-center bg-black bg-opacity-50">
                    <View className="w-10/12 lg:w-1/4 p-6 bg-[#1a1a1a] rounded-lg">
                        <Text className="text-white text-xl text-center mb-4">{feedbackMessage}</Text>
                        <TouchableOpacity
                            className="bg-red-500 p-3 rounded-lg"
                            onPress={() => {
                                setFeedbackModalVisible(false);
                                router.push("/home");
                            }}
                        >
                            <Text className="text-white text-center">OK</Text>
                        </TouchableOpacity>
                    </View>
                </View>
            </Modal>

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

export default AppointmentScreen;
