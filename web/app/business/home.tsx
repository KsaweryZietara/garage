import React, {useState, useEffect} from "react";
import {
    View,
    Text,
    Platform,
    ActivityIndicator,
    FlatList, StatusBar,
} from "react-native";
import axios from "axios";
import {get, remove} from "@/utils/auth";
import moment, {Moment} from "moment";
import "moment/locale/pl";
import CalendarStrip from "react-native-calendar-strip";
import {getJwtPayload} from "@/utils/jwt";
import {useRouter} from "expo-router";
import {EmployeeAppointment} from "@/types";
import {EMPLOYEE_JWT} from "@/constants/constants";
import {formatTime} from "@/utils/time";
import BusinessMenu from "@/components/BusinessMenu";
import CustomButton from "@/components/CustomButton";

moment.locale("pl");

const HomeScreen = () => {
    const router = useRouter();
    const [email, setEmail] = useState<string | null>(null);
    const [role, setRole] = useState<string | null>(null);
    const [menuVisible, setMenuVisible] = useState(false);
    const [garageName, setGarageName] = useState("");
    const [appointments, setAppointments] = useState<EmployeeAppointment[]>([]);
    const [loadingAppointments, setLoadingAppointments] = useState<boolean>(true);
    const [selectedDate, setSelectedDate] = useState<Moment>(moment());

    useEffect(() => {
        const fetchGarageName = async () => {
            const token = await get(EMPLOYEE_JWT);
            await axios.get("/api/employees/garages", {
                headers: {"Authorization": `Bearer ${token}`}
            })
                .then((response) => {
                    if (response.data && response.data.name) {
                        setGarageName(response.data.name);
                    }
                })
                .catch((error) => {
                    console.error(error)
                });
        };

        const fetchJwtPayload = async () => {
            const jwtPayload = await getJwtPayload(EMPLOYEE_JWT);
            setEmail(jwtPayload?.email || null);
            setRole(jwtPayload?.role || null)
        };

        fetchGarageName();
        fetchJwtPayload();
    }, []);

    useEffect(() => {
        fetchAppointments();
    }, [selectedDate]);

    const fetchAppointments = async () => {
        setLoadingAppointments(true);
        const token = await get(EMPLOYEE_JWT);
        await axios.get<EmployeeAppointment[]>(`/api/employees/appointments?date=${selectedDate.format("YYYY-MM-DD")}`, {
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

    const handleAppointmentDelete = async (id: number) => {
        const token = await get(EMPLOYEE_JWT);
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

    const handleDateChange = (date: Moment) => {
        setSelectedDate(date);
    };

    const renderAppointmentItem = ({item}: { item: EmployeeAppointment }) => {
        const startHour = new Date(item.startTime).getUTCHours();
        const timeUntilStart = new Date(item.startTime).getTime() - new Date().getTime();
        const isMoreThan24Hours = timeUntilStart > 24 * 60 * 60 * 1000;

        return (
            <View className="p-4 my-2 mx-4 bg-gray-700 rounded-lg flex-col lg:flex-row justify-between lg:items-center">
                <View>
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
                    <Text className="text-sm text-white">
                        Samochód: {item.car.make} {item.car.model}
                    </Text>
                </View>
                {isMoreThan24Hours && (
                    <CustomButton
                        title={"Anuluj wizytę"}
                        onPress={() => {
                            handleAppointmentDelete(item.id)
                        }}
                        containerStyles="bg-white self-center mt-3 lg:mt-0 w-3/5 lg:w-1/5"
                        textStyles="text-gray-700 font-bold"
                    />
                )}
            </View>
        );
    };

    return (
        <View className="flex-1">
            <View className="flex-row justify-between p-4 bg-gray-700">
                <Text className="text-2xl lg:text-4xl font-bold text-white lg:mt-1.5" onPress={() => {
                    router.push("/business/home")
                }}>
                    {garageName.toUpperCase()}
                </Text>
                <Text
                    className="text-white font-bold lg:text-xl"
                    onPress={() => setMenuVisible(true)}
                    style={{
                        borderRadius: 5,
                        padding: Platform.OS === "web" ? 12 : 6,
                        marginRight: 5,
                    }}
                >
                    {email}
                </Text>
            </View>
            <CalendarStrip
                scrollable
                locale={{name: "pl", config: {}}}
                calendarAnimation={{type: "sequence", duration: 30}}
                style={{height: 100, paddingBottom: Platform.OS === "web" ? 100 : 10}}
                calendarHeaderStyle={{
                    color: "white",
                    marginBottom: Platform.OS === "web" ? 40 : 10,
                    fontSize: 20,
                }}
                calendarColor={"#374151"}
                dateNumberStyle={{color: "white"}}
                dateNameStyle={{color: "white"}}
                highlightDateNumberStyle={{color: "#111827"}}
                highlightDateNameStyle={{color: "#111827"}}
                disabledDateNameStyle={{color: "gray"}}
                disabledDateNumberStyle={{color: "gray"}}
                iconContainer={{flex: 0.1}}
                markedDates={[
                    {
                        date: moment().format("YYYY-MM-DD"),
                        dots: [{color: "#111827", selectedColor: "#111827"}],
                    },
                ]}
                leftSelector={<Text style={{color: "white", fontSize: 30}}>&lt;</Text>}
                rightSelector={<Text style={{color: "white", fontSize: 30}}>&gt;</Text>}
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

            <BusinessMenu
                menuVisible={menuVisible}
                onClose={() => setMenuVisible(false)}
                role={role}
                email={email}
                onLogout={() => {
                    remove(EMPLOYEE_JWT);
                    setEmail(null);
                    setMenuVisible(false)
                    router.push("/business/login")
                }}
            />
            <StatusBar backgroundColor="#374151"/>
        </View>
    );
};

export default HomeScreen;
