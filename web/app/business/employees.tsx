import {useRouter} from "expo-router";
import React, {useEffect, useState} from "react";
import {get, remove} from "@/utils/auth";
import {EMPLOYEE_JWT} from "@/constants/constants";
import axios from "axios";
import {getJwtPayload} from "@/utils/jwt";
import {
    ActivityIndicator, FlatList, Modal,
    Platform, StatusBar,
    Text, TextInput, TouchableWithoutFeedback,
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
    const [createEmployeeVisible, setCreateEmployeeVisible] = useState<boolean>(false);
    const [employeeEmail, setEmployeeEmail] = useState<string>("");
    const [errorMessage, setErrorMessage] = useState("");
    const [resendEmailModalVisible, setResendEmailModalVisible] = useState<boolean>(false);
    const [selectedEmployeeId, setSelectedEmployeeId] = useState<number | null>(null);

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

    const handleResendEmail = async () => {
        if (!selectedEmployeeId) {
            return;
        }
        setResendEmailModalVisible(false);
        setSelectedEmployeeId(null);
        const token = await get(EMPLOYEE_JWT);
        await axios.get(`/api/employees/${selectedEmployeeId}/confirmation`, {
            headers: {"Authorization": `Bearer ${token}`}
        })
            .catch((error) => {
                console.error(error);
            });
    };

    const validateEmployee = () => {
        if (!employeeEmail.trim()) {
            return "Adres e-mail musi być wypełniony.";
        }

        if (employeeEmail.length > 255) {
            return "Adres e-mail nie może przekraczać 255 znaków.";
        }

        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        if (!emailRegex.test(employeeEmail)) {
            return "Nieprawidłowy format adresu e-mail.";
        }

        return null;
    };

    const handleCreateEmployee = async () => {
        const validationError = validateEmployee();
        if (validationError) {
            setErrorMessage(validationError);
            return;
        }
        const token = await get(EMPLOYEE_JWT);
        const data = {
            email: employeeEmail
        };
        await axios.post("/api/employees", data, {headers: {"Authorization": `Bearer ${token}`}})
            .then(() => {
                setErrorMessage("");
                setEmployeeEmail("")
                setCreateEmployeeVisible(false)
                fetchEmployees()
            })
            .catch((error) => {
                console.error(error)
                setErrorMessage(error.response.data.message);
            });
    };

    const handleDeleteEmployee = async (id: number) => {
        const token = await get(EMPLOYEE_JWT);
        await axios.delete(`/api/employees/${id}`, {headers: {"Authorization": `Bearer ${token}`}})
            .then(() => {
                fetchEmployees()
            })
            .catch((error) => {
                console.error(error)
            });
    }

    const renderEmployeeItem = ({item}: { item: Employee }) => (
        <View className="p-4 my-2 mx-4 bg-gray-700 rounded-lg">
            <View className="flex-col lg:flex-row justify-between lg:items-center">
                <View className="flex-1">
                    <Text className="text-2xl text-white font-bold">
                        {item.email}
                    </Text>
                    {item.confirmed ? (
                        <View>
                            <Text className="text-lg text-white font-bold">
                                {item.name} {item.surname}
                            </Text>
                            <Text className={"text-sm text-green-500"}>
                                Zarejestrowany
                            </Text>
                        </View>
                    ) : (
                        <View>
                            <Text className={"text-sm text-red-500 underline"} onPress={() => {
                                setSelectedEmployeeId(item.id);
                                setResendEmailModalVisible(true);
                            }}>
                                Nie zarejestrowany
                            </Text>
                        </View>
                    )}
                </View>

                <CustomButton
                    title={"Usuń pracownika"}
                    onPress={() => {
                        handleDeleteEmployee(item.id)
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
                    <Text className="text-gray-700 text-3xl font-bold mt-5 ml-4 mb-3">Pracownicy</Text>
                </View>

                <CustomButton
                    title={"Dodaj pracownika"}
                    onPress={() => {
                        setEmployeeEmail("")
                        setErrorMessage("")
                        setCreateEmployeeVisible(true)
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

            <Modal
                visible={resendEmailModalVisible}
                animationType="fade"
                transparent={true}
                onRequestClose={() => setResendEmailModalVisible(false)}
            >
                <TouchableWithoutFeedback onPress={() => setResendEmailModalVisible(false)}>
                    <View className="flex-1 justify-center items-center"
                          style={{backgroundColor: 'rgba(0, 0, 0, 0.75)'}}>
                        <TouchableWithoutFeedback onPress={() => {
                        }}>
                            <View className="bg-gray-700 py-5 rounded-lg w-4/5 lg:w-2/5">
                                <View>
                                    <Text className="text-white text-lg font-bold self-center">
                                        Czy na pewno chcesz wysłać ponownie e-mail potwierdzający?
                                    </Text>
                                    <CustomButton
                                        title="Wyślij"
                                        onPress={handleResendEmail}
                                        containerStyles="bg-white self-center mt-5 px-10 lg:px-20"
                                        textStyles="text-gray-700 font-bold"
                                    />
                                </View>
                            </View>
                        </TouchableWithoutFeedback>
                    </View>
                </TouchableWithoutFeedback>
            </Modal>

            <Modal
                visible={createEmployeeVisible}
                animationType="fade"
                transparent={true}
                onRequestClose={() => setCreateEmployeeVisible(false)}
            >
                <TouchableWithoutFeedback onPress={() => setCreateEmployeeVisible(false)}>
                    <View className="flex-1 justify-center items-center"
                          style={{backgroundColor: 'rgba(0, 0, 0, 0.75)'}}>
                        <TouchableWithoutFeedback onPress={() => {
                        }}>
                            <View className="bg-gray-700 p-5 rounded-lg w-4/5 lg:w-2/5">
                                <View>
                                    <Text className="text-white text-lg font-bold">Dodaj pracownika</Text>
                                    <TextInput
                                        value={employeeEmail}
                                        onChangeText={setEmployeeEmail}
                                        placeholder="Adres email pracownika"
                                        className="border p-2 mt-5 rounded text-#2d2d2d bg-white align-text-top max-h-20"
                                        placeholderTextColor="#2d2d2d"
                                    />
                                    {errorMessage && (
                                        <Text className="text-red-500 text-center mt-3">
                                            {errorMessage}
                                        </Text>
                                    )}
                                    <CustomButton
                                        title="Dodaj"
                                        onPress={handleCreateEmployee}
                                        containerStyles="bg-white self-center mt-5 px-10 lg:px-20"
                                        textStyles="text-gray-700 font-bold"
                                    />
                                </View>
                            </View>
                        </TouchableWithoutFeedback>
                    </View>
                </TouchableWithoutFeedback>
            </Modal>

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
