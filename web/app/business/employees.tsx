import {useRouter} from "expo-router";
import React, {useEffect, useState} from "react";
import {get, remove} from "@/utils/auth";
import {EMPLOYEE_JWT} from "@/constants/constants";
import axios from "axios";
import {getJwtPayload} from "@/utils/jwt";
import {
    ActivityIndicator, FlatList,
    Platform, StatusBar,
    Text,
    View
} from "react-native";
import BusinessMenu from "@/components/BusinessMenu";
import {Employee, Garage} from "@/types";
import CustomButton from "@/components/CustomButton";

const EmployeesScreen = () => {
    const router = useRouter();
    const [email, setEmail] = useState<string | null>(null);
    const [role, setRole] = useState<string | null>(null);
    const [menuVisible, setMenuVisible] = useState(false);
    const [garage, setGarage] = useState<Garage | null>(null);
    const [employees, setEmployees] = useState<Employee[]>([]);
    const [loadingEmployees, setLoadingEmployees] = useState<boolean>(false);

    useEffect(() => {
        const fetchGarageName = async () => {
            const token = await get(EMPLOYEE_JWT);
            await axios.get<Garage>("/api/employees/garages", {
                headers: {"Authorization": `Bearer ${token}`}
            })
                .then((response) => {
                    if (response.data) {
                        setGarage(response.data);
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
        fetchEmployees()
    }, [garage]);

    const fetchEmployees = async () => {
        setLoadingEmployees(true);
        const token = await get(EMPLOYEE_JWT);
        await axios.get<Employee[]>("/api/employees", {
            headers: {"Authorization": `Bearer ${token}`}
        })
            .then((response) => {
                if (response.data) {
                    setEmployees(response.data);
                }
            })
            .catch((error) => {
                console.error(error);
            }).finally(() => {
                setLoadingEmployees(false);
            });
    };

    const renderEmployeeItem = ({item}: { item: Employee }) => (
        <View className="p-4 my-2 mx-4 bg-gray-700 rounded-lg">
            <View className="flex-col lg:flex-row justify-between lg:items-center">
                <View className="flex-1">
                    <Text className="text-lg text-white font-bold">
                        {item.name} {item.surname}
                    </Text>
                    {item.confirmed ? (
                        <Text className={"text-sm text-green-500"}>
                            Zarejestrowany
                        </Text>
                    ) : (
                        <View>
                            <Text className={"text-sm text-red-500 underline"} onPress={() => {
                            }}>
                                Nie zarejestrowany
                            </Text>
                        </View>
                    )}
                </View>

                <CustomButton
                    title={"Usuń pracownika"}
                    onPress={() => {
                    }}
                    containerStyles="bg-white self-center mt-3 lg:mt-0 w-3/5 lg:w-1/5"
                    textStyles="text-gray-700 font-bold"
                />
            </View>
        </View>
    );

    return (
        <View className="flex-1">
            <View className="flex-row justify-between p-4 bg-gray-700">
                <Text className="text-lg lg:text-4xl font-bold text-white">
                    {(garage?.name ? garage.name.toUpperCase() : "GARAGE")}
                </Text>
                <Text
                    className="text-white font-bold"
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

            <View className="flex-row justify-between lg:items-center">
                <View className="flex-1">
                    <Text className="text-gray-700 text-3xl font-bold mt-5 ml-4 mb-3">Usługi</Text>
                </View>

                <CustomButton
                    title={"Dodaj pracownika"}
                    onPress={() => {
                    }}
                    containerStyles="bg-gray-700 self-center mt-3 w-2/5 lg:w-1/5 mr-3"
                    textStyles="text-white font-bold"
                />
            </View>

            {loadingEmployees ? (
                <ActivityIndicator size="large" color="#374151"/>
            ) : (
                <FlatList
                    data={employees}
                    renderItem={renderEmployeeItem}
                    keyExtractor={(item, index) => item?.id?.toString() || index.toString()}
                    ListEmptyComponent={
                        <View className="flex-1 justify-center items-center mb-40">
                            <Text className="text-gray-700 text-2xl mt-10">Brak pracowników.</Text>
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

export default EmployeesScreen;
