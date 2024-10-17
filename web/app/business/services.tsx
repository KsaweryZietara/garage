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
import {Garage, Service} from "@/types";
import CustomButton from "@/components/CustomButton";

const ServicesScreen = () => {
    const router = useRouter();
    const [email, setEmail] = useState<string | null>(null);
    const [role, setRole] = useState<string | null>(null);
    const [menuVisible, setMenuVisible] = useState(false);
    const [garage, setGarage] = useState<Garage | null>(null);
    const [services, setServices] = useState<Service[]>([]);
    const [loadingServices, setLoadingServices] = useState<boolean>(false);
    const [serviceVisible, setServiceVisible] = useState<boolean>(false);
    const [name, setName] = useState<string>("");
    const [time, setTime] = useState<string>("");
    const [price, setPrice] = useState<string>("");
    const [errorMessage, setErrorMessage] = useState("");

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
        fetchServices()
    }, [garage]);

    const fetchServices = async () => {
        setLoadingServices(true);
        await axios.get<Service[]>(`/api/garages/${garage?.id}/services`)
            .then((response) => {
                if (response.data) {
                    setServices(response.data);
                }
            })
            .catch((error) => {
                console.error(error);
            }).finally(() => {
                setLoadingServices(false);
            });
    };

    const validateService = () => {
        if (!name || !time || !price) {
            return "Wszystkie pola dla usługi muszą być wypełnione.";
        }

        if (name.length > 255) {
            return "Nazwa nie może przekraczać 255 znaków.";
        }

        const timeNumber = Number(time);
        const priceNumber = Number(price);

        if (isNaN(timeNumber) || timeNumber <= 0 || !Number.isInteger(timeNumber)) {
            return "Czas musi być podany w pełnych godzinach.";
        }

        if (timeNumber > 720) {
            return "Czas usługi nie może być dłuższy niż miesiąc.";
        }

        if (isNaN(priceNumber) || priceNumber <= 0 || !Number.isInteger(priceNumber)) {
            return "Cena musi być podana w pełnych złotówkach.";
        }

        return null;
    };

    const handleCreateService = async () => {
        const validationError = validateService();
        if (validationError) {
            setErrorMessage(validationError);
            return;
        }
        const token = await get(EMPLOYEE_JWT);
        const data = {
            name: name,
            time: parseInt(time, 10),
            price: parseInt(price, 10)
        };
        await axios.post("/api/services", data, {headers: {"Authorization": `Bearer ${token}`}})
            .then(() => {
                setErrorMessage("");
                setName("")
                setTime("")
                setPrice("")
                setServiceVisible(false)
                fetchServices()
            })
            .catch((error) => {
                console.error(error)
                setErrorMessage(error.response.data.message);
            });
    };

    const handleDeleteService = async (Id: number) => {
        const token = await get(EMPLOYEE_JWT);
        await axios.delete(`/api/services/${Id}`, {
            headers: {"Authorization": `Bearer ${token}`}
        })
            .then(() => {
                fetchServices();
            })
            .catch((error) => {
                console.error(error);
            });
    };

    const renderServiceItem = ({item}: { item: Service }) => (
        <View className="p-4 my-2 mx-4 bg-gray-700 rounded-lg">
            <View className="flex-col lg:flex-row justify-between lg:items-center">
                <View className="flex-1">
                    <Text className="text-lg text-white font-bold">
                        {item.name}
                    </Text>
                    <Text className="text-sm text-[#ddd]">
                        Czas: {item.time}
                    </Text>
                    <Text className="text-sm text-[#ddd]">
                        Cena: {item.price}
                    </Text>
                </View>

                <CustomButton
                    title={"Usuń usługę"}
                    onPress={() => {
                        handleDeleteService(item.id)
                    }}
                    containerStyles="bg-white self-center mt-3 lg:mt-0 w-2/5 lg:w-1/5"
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
                    title={"Dodaj usługę"}
                    onPress={() => {
                        setServiceVisible(true);
                        setName("");
                        setTime("");
                        setPrice("")
                        setErrorMessage("")
                    }}
                    containerStyles="bg-gray-700 self-center mt-3 w-2/5 lg:w-1/5 mr-3"
                    textStyles="text-white font-bold"
                />
            </View>

            {loadingServices ? (
                <ActivityIndicator size="large" color="#374151"/>
            ) : (
                <FlatList
                    data={services}
                    renderItem={renderServiceItem}
                    keyExtractor={(item, index) => item?.id?.toString() || index.toString()}
                    ListEmptyComponent={
                        <View className="flex-1 justify-center items-center mb-40">
                            <Text className="text-gray-700 text-2xl mt-10">Brak usług.</Text>
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

            <Modal
                visible={serviceVisible}
                animationType="fade"
                transparent={true}
                onRequestClose={() => setServiceVisible(false)}
            >
                <TouchableWithoutFeedback onPress={() => setServiceVisible(false)}>
                    <View className="flex-1 justify-center items-center"
                          style={{backgroundColor: 'rgba(0, 0, 0, 0.75)'}}>
                        <TouchableWithoutFeedback onPress={() => {
                        }}>
                            <View className="bg-gray-700 p-5 rounded-lg w-4/5 lg:w-2/5">
                                <View>
                                    <Text className="text-white text-lg font-bold">Dodaj usługę</Text>
                                    <TextInput
                                        value={name}
                                        onChangeText={setName}
                                        placeholder="Nazwa usługi"
                                        className="border p-2 mt-5 rounded text-#2d2d2d bg-white align-text-top max-h-20"
                                        placeholderTextColor="#2d2d2d"
                                    />
                                    <TextInput
                                        value={time}
                                        onChangeText={setTime}
                                        placeholder="Czas wykonania (w godzinach)"
                                        className="border p-2 mt-5 rounded text-#2d2d2d bg-white align-text-top max-h-20"
                                        placeholderTextColor="#2d2d2d"
                                    />
                                    <TextInput
                                        value={price}
                                        onChangeText={setPrice}
                                        placeholder="Cena (zł)"
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
                                        onPress={handleCreateService}
                                        containerStyles="bg-white self-center mt-5 px-10 lg:px-20"
                                        textStyles="text-gray-700 font-bold"
                                    />
                                </View>
                            </View>
                        </TouchableWithoutFeedback>
                    </View>
                </TouchableWithoutFeedback>
            </Modal>
            <StatusBar backgroundColor="#374151"/>
        </View>
    );
};

export default ServicesScreen;
