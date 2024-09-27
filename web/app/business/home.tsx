import React, {useState, useEffect} from "react";
import {View, Text, Platform, ActivityIndicator, FlatList} from "react-native";
import axios from "axios";
import {getJWT} from "@/utils/auth";
import moment, {Moment} from "moment";
import 'moment/locale/pl';
import CalendarStrip from "react-native-calendar-strip";

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
    const [garageName, setGarageName] = useState("garage");
    const [appointments, setAppointments] = useState<Appointment[]>([]);
    const [loadingAppointments, sideloadingAppointments] = useState<boolean>(true);
    const [selectedDate, setSelectedDate] = useState<Moment>(moment());

    useEffect(() => {
        const fetchGarageName = async () => {
            try {
                const token = await getJWT();
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

        fetchGarageName();
    }, []);

    useEffect(() => {
        const fetchAvailableSlots = async () => {
            sideloadingAppointments(true);
            try {
                const token = await getJWT();
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
            <View className="flex-row justify-start p-4 bg-gray-700">
                <Text className="text-lg lg:text-4xl font-bold text-white">
                    {garageName.toUpperCase()}
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
        </View>
    );
};

export default HomeScreen;
