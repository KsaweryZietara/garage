import {useLocalSearchParams, useRouter} from "expo-router";
import React, {useEffect, useState} from "react";
import {ActivityIndicator, FlatList, Modal, Platform, StatusBar, Text, TouchableOpacity, View} from "react-native";
import CalendarStrip from "react-native-calendar-strip";
import axios from "axios";
import moment, {Moment} from "moment";
import "moment/locale/pl";
import {get} from "@/utils/auth";
import {getJwtPayload} from "@/utils/jwt";
import EmailDisplay from "@/components/EmailDisplay";
import MenuModal from "@/components/MenuModal";
import {Employee, Make, Model, Service, TimeSlot} from "@/types";
import {CUSTOMER_JWT} from "@/constants/constants";
import {formatDateTime, formatTime} from "@/utils/time";
import {Picker} from "@react-native-picker/picker";

moment.locale("pl");

const AppointmentScreen = () => {
    const router = useRouter();
    const [email, setEmail] = useState<string | null>(null);
    const [menuVisible, setMenuVisible] = useState(false);
    const {serviceid, employeeid} = useLocalSearchParams() as { serviceid: string, employeeid: string };
    const [service, setService] = useState<Service | null>(null);
    const [employee, setEmployee] = useState<Employee | null>(null);
    const [timeSlots, setTimeSlots] = useState<TimeSlot[]>([]);
    const [makes, setMakes] = useState<Make[]>([]);
    const [selectedMake, setSelectedMake] = useState<number | undefined>(undefined);
    const [models, setModels] = useState<Model[]>([]);
    const [selectedModel, setSelectedModel] = useState<number | undefined>(undefined);
    const [loadingService, setLoadingService] = useState<boolean>(true);
    const [loadingEmployee, setLoadingEmployee] = useState<boolean>(true);
    const [loadingSlots, setLoadingSlots] = useState<boolean>(true);
    const [selectedDate, setSelectedDate] = useState<Moment>(moment());
    const [modalVisible, setModalVisible] = useState<boolean>(false);
    const [selectedTimeSlot, setSelectedTimeSlot] = useState<TimeSlot | null>(null);
    const [feedbackModalVisible, setFeedbackModalVisible] = useState<boolean>(false);
    const [feedbackMessage, setFeedbackMessage] = useState<string>("");
    const [errorMessage, setErrorMessage] = useState("");

    useEffect(() => {
        const fetchEmail = async () => {
            const jwtPayload = await getJwtPayload(CUSTOMER_JWT);
            setEmail(jwtPayload?.email || null);
        };

        const fetchMakes = async () => {
            await axios.get<Make[]>(
                "/api/makes"
            )
                .then((response) => {
                    setMakes(response.data);
                })
                .catch((error) => {
                    console.error(error);
                });
        };

        fetchEmail();
        fetchMakes()
    }, []);

    useEffect(() => {
        setSelectedModel(undefined)
        const fetchModels = async () => {
            await axios.get<Model[]>(
                `/api/makes/${selectedMake}/models`
            )
                .then((response) => {
                    setModels(response.data);
                })
                .catch((error) => {
                    console.error(error);
                });
        };

        fetchModels()
    }, [selectedMake]);

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
        if (!selectedModel) {
            setErrorMessage("Wybierz model samochodu.")
            return;
        }
        setErrorMessage("")
        setModalVisible(false);
        const token = await get(CUSTOMER_JWT);
        if (token == null) {
            router.push("/login")
            return
        }
        const data = {
            modelId: parseInt(String(selectedModel), 10),
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
                setErrorMessage("")
                setSelectedMake(undefined)
                setModalVisible(true);
            }}
            className="p-4 my-2 mx-4 bg-[#2d2d2d] rounded-lg"
        >
            <Text className="text-lg font-bold text-white">
                {formatTime(item.startTime)}
            </Text>
            <Text className="text-md text-gray-400">
                {formatDateTime(item.endTime)}
            </Text>
        </TouchableOpacity>
    );

    return (
        <View className="flex-1 bg-black">
            <View className="flex-row justify-between p-4 bg-black">
                <Text className="text-white text-2xl lg:text-4xl font-bold lg:mt-1.5"
                      onPress={() => router.push("/home")}>GARAGE</Text>
                <EmailDisplay email={email} setMenuVisible={setMenuVisible}/>
            </View>

            {loadingService || loadingEmployee ? (
                <ActivityIndicator size="large" color="#ff5c5c"/>
            ) : (
                service && (
                    <View className="p-6 bg-[#1a1a1a] rounded-lg mx-4 mt-4 shadow-lg">
                        <Text className="text-3xl font-extrabold text-white mb-2">{service.name}</Text>
                        <Text className="text-xl text-[#aaa]">Czas: {service.time} godz.</Text>
                        <Text className="text-xl text-[#aaa]">Cena: {service.price} zł</Text>
                        <Text className="text-xl text-[#aaa]">Mechanik: {employee?.name} {employee?.surname}</Text>
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
                                    <Text className="text-white text-xl mb-2">{service.price} zł</Text>
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
                                <View className="text-white text-xl mb-2 flex-row justify-between">
                                    <Text className="text-white text-xl mb-2">Odbiór:</Text>
                                    <Text className="text-white text-xl mb-2">
                                        {formatDateTime(selectedTimeSlot.endTime)}
                                    </Text>
                                </View>
                                <View className="mb-2">
                                    <Text className="text-white text-xl mb-2">Wybierz markę:</Text>
                                    <Picker
                                        selectedValue={selectedMake}
                                        onValueChange={(itemValue) => setSelectedMake(itemValue)}
                                        style={{
                                            color: 'black',
                                            fontSize: 15,
                                            height: Platform.OS === 'web' ? 30 : undefined,
                                            backgroundColor: 'white'
                                        }}
                                    >
                                        <Picker.Item label="Wybierz markę" value={null}/>
                                        {makes.map(make => (
                                            <Picker.Item key={make.id} label={make.name} value={make.id}/>
                                        ))}
                                    </Picker>
                                </View>
                                <View className="mb-2">
                                    <Text className="text-white text-xl mb-2">Wybierz model:</Text>
                                    <Picker
                                        selectedValue={selectedModel}
                                        onValueChange={(itemValue) => setSelectedModel(itemValue)}
                                        style={{
                                            color: 'black',
                                            fontSize: 15,
                                            height: Platform.OS === 'web' ? 30 : undefined,
                                            backgroundColor: 'white'
                                        }}
                                    >
                                        <Picker.Item label="Wybierz model" value={null}/>
                                        {models.map(make => (
                                            <Picker.Item key={make.id} label={make.name} value={make.id}/>
                                        ))}
                                    </Picker>
                                </View>
                                {errorMessage && (
                                    <Text className="text-red-500 text-center mt-2">
                                        {errorMessage}
                                    </Text>
                                )}
                                <TouchableOpacity
                                    className="bg-red-500 p-3 rounded-lg mt-5 mb-4"
                                    onPress={handleSubmit}
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
