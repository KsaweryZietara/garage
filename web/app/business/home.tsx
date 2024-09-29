import React, {useState, useEffect} from "react";
import {
    View,
    Text,
    Platform,
    ActivityIndicator,
    FlatList,
    Modal,
    TouchableWithoutFeedback,
    TouchableOpacity
} from "react-native";
import axios from "axios";
import {get, remove} from "@/utils/auth";
import moment, {Moment} from "moment";
import 'moment/locale/pl';
import CalendarStrip from "react-native-calendar-strip";
import {getEmail} from "@/utils/jwt";
import {useRouter} from "expo-router";

moment.locale('pl');

interface Appointment {
    id: number;
    startTime: Date;
    endTime: Date;
    service: Service;
    employee?: Employee;
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

const HomeScreen = () => {
    const router = useRouter();
    const [email, setEmail] = useState<string | null>(null);
    const [menuVisible, setMenuVisible] = useState(false);
    const [garageName, setGarageName] = useState("garage");
    const [appointments, setAppointments] = useState<Appointment[]>([]);
    const [loadingAppointments, sideloadingAppointments] = useState<boolean>(true);
    const [selectedDate, setSelectedDate] = useState<Moment>(moment());

    useEffect(() => {
        const fetchGarageName = async () => {
            try {
                const token = await get("employee_jwt");
                const response = await axios.get("/api/employee/garage", {
                    headers: {"Authorization": `Bearer ${token}`}
                });
                if (response.data && response.data.name) {
                    setGarageName(response.data.name);
                }
            } catch (error) {
                console.log(error);
            }
        };

        const fetchEmail = async () => {
            const email = await getEmail("employee_jwt");
            setEmail(email);
        };

        fetchGarageName();
        fetchEmail();
    }, []);

    useEffect(() => {
        const fetchAvailableSlots = async () => {
            sideloadingAppointments(true);
            try {
                const token = await get("employee_jwt");
                const response = await axios.get<Appointment[]>(`/api/employee/appointments?date=${selectedDate.format("YYYY-MM-DD")}`, {
                    headers: {"Authorization": `Bearer ${token}`}
                });
                setAppointments(response.data);
            } catch (error) {
                console.error(error);
            } finally {
                sideloadingAppointments(false);
            }
        };

        fetchAvailableSlots();
    }, [selectedDate]);

    const handleLogout = () => {
        remove("employee_jwt");
        setEmail(null);
        setMenuVisible(false)
        router.push("/business/login")
    };

    const handleMenuClose = () => {
        setMenuVisible(false)
    }

    const handleDateChange = (date: Moment) => {
        setSelectedDate(date);
    };

    const formatTime = (date: Date): string => {
        const d = new Date(date);
        const hours = d.getUTCHours().toString().padStart(2, '0');
        const minutes = d.getUTCMinutes().toString().padStart(2, '0');
        return `${hours}:${minutes}`;
    };

    const renderAppointmentItem = ({item}: { item: Appointment }) => {
        const startHour = new Date(item.startTime).getUTCHours();
        console.log(startHour)

        return (
            <View className="p-4 my-2 mx-4 bg-gray-700 rounded-lg">
                <Text className="text-lg font-bold text-white">
                    {startHour === 16
                        ? `${formatTime(item.startTime)}`
                        : `${formatTime(item.startTime)} - ${formatTime(item.endTime)}`}
                </Text>
                {item.employee && (
                    <Text className="text-sm text-white">
                        Mechanik: {item.employee.name} {item.employee.surname}
                    </Text>
                )}
                <Text className="text-sm text-white">
                    Usługa: {startHour === 16 ? "Przyjęcie samochodu" : item.service.name}
                </Text>
            </View>
        );
    };

    return (
        <View className="flex-1">
            <View className="flex-row justify-between p-4 bg-gray-700">
                <Text className="text-lg lg:text-4xl font-bold text-white">
                    {garageName.toUpperCase()}
                </Text>
                <Text
                    className="text-white font-bold"
                    onPress={() => setMenuVisible(true)}
                    style={{
                        borderRadius: 5,
                        padding: Platform.OS === 'web' ? 12 : 6,
                        marginRight: 5,
                    }}
                >
                    {email}
                </Text>
            </View>
            <CalendarStrip
                scrollable
                locale={{name: 'pl', config: {}}}
                calendarAnimation={{type: 'sequence', duration: 30}}
                style={{height: 100, paddingBottom: Platform.OS === 'web' ? 100 : 10}}
                calendarHeaderStyle={{
                    color: 'white',
                    marginBottom: Platform.OS === 'web' ? 40 : 10,
                    fontSize: 20,
                }}
                calendarColor={'#374151'}
                dateNumberStyle={{color: 'white'}}
                dateNameStyle={{color: 'white'}}
                highlightDateNumberStyle={{color: '#111827'}}
                highlightDateNameStyle={{color: '#111827'}}
                disabledDateNameStyle={{color: 'gray'}}
                disabledDateNumberStyle={{color: 'gray'}}
                iconContainer={{flex: 0.1}}
                markedDates={[
                    {
                        date: moment().format('YYYY-MM-DD'),
                        dots: [{color: '#111827', selectedColor: '#111827'}],
                    },
                ]}
                leftSelector={<Text style={{color: 'white', fontSize: 30}}>&lt;</Text>}
                rightSelector={<Text style={{color: 'white', fontSize: 30}}>&gt;</Text>}
                onDateSelected={handleDateChange}
                selectedDate={selectedDate}
            />
            <Text className="text-gray-700 text-2xl font-bold mt-4 ml-4 mb-3">Wizyty</Text>

            {loadingAppointments ? (
                <ActivityIndicator size="large" color="#374151"/>
            ) : (
                <FlatList
                    data={appointments}
                    renderItem={renderAppointmentItem}
                    keyExtractor={(item, index) => index.toString()}
                    ListEmptyComponent={
                        <View className="flex-1 justify-center items-center mb-40">
                            <Text className="text-gray-700 text-xl">Brak wizyt na dziś.</Text>
                        </View>
                    }
                    showsHorizontalScrollIndicator={false}
                />
            )}

            <Modal
                transparent={true}
                animationType="fade"
                visible={menuVisible}
                onRequestClose={handleMenuClose}
            >
                <TouchableWithoutFeedback onPress={handleMenuClose}>
                    <View style={{
                        flex: 1,
                        justifyContent: 'flex-start',
                        alignItems: 'flex-end',
                    }}>
                        <View style={{
                            marginRight: Platform.OS === 'web' ? 32 : 27,
                            marginTop: Platform.OS === 'web' ? 52 : 50,
                            backgroundColor: 'white',
                            borderRadius: 5,
                            padding: Platform.OS === 'web' ? 12 : 6,
                            elevation: 5,
                            shadowColor: '#000',
                            shadowOffset: {
                                width: 0,
                                height: 2,
                            },
                            shadowOpacity: 0.25,
                            shadowRadius: 4,
                        }}>
                            {email && (
                                <TouchableOpacity onPress={handleLogout}>
                                    <Text className="text-red-500 font-bold">Wyloguj</Text>
                                </TouchableOpacity>
                            )}
                        </View>
                    </View>
                </TouchableWithoutFeedback>
            </Modal>
        </View>
    );
};

export default HomeScreen;
