import {useRouter} from "expo-router";
import React, {useEffect, useState} from "react";
import {get, remove} from "@/utils/auth";
import {EMPLOYEE_JWT} from "@/constants/constants";
import axios from "axios";
import {getJwtPayload} from "@/utils/jwt";
import {
    Platform, StatusBar,
    Text,
    View,
    Image,
    Dimensions, ActivityIndicator, TouchableWithoutFeedback, Modal
} from "react-native";
import BusinessMenu from "@/components/BusinessMenu";
import {Garage} from "@/types";
import * as ImagePicker from 'expo-image-picker';
import * as FileSystem from 'expo-file-system';
import CustomButton from "@/components/CustomButton";

const LogoScreen = () => {
    const router = useRouter();
    const {width, height} = Dimensions.get("window");
    const [email, setEmail] = useState<string | null>(null);
    const [role, setRole] = useState<string | null>(null);
    const [menuVisible, setMenuVisible] = useState(false);
    const [garage, setGarage] = useState<Garage | null>(null);
    const [loading, setLoading] = useState(false);
    const [image, setImage] = useState<string | null>(null);
    const [errorMessage, setErrorMessage] = useState("");
    const [logoSubmitted, setLogoSubmitted] = useState<boolean>(false);

    const handlePickLogo = async () => {
        setErrorMessage("")
        let result = await ImagePicker.launchImageLibraryAsync({
            mediaTypes: ImagePicker.MediaTypeOptions.All,
            allowsEditing: true,
            aspect: [4, 3],
            quality: 1,
        });

        if (!result.canceled) {
            setImage(result.assets[0].uri);
        }
    };

    const handleSaveLogo = async () => {
        setErrorMessage("")
        if (!image) {
            setErrorMessage("Wybierz logo.")
            return;
        }

        setLoading(true);
        try {
            let base64Image;
            if (Platform.OS === "web") {
                const response = await fetch(image);
                const blob = await response.blob();
                base64Image = await new Promise((resolve, reject) => {
                    const reader = new FileReader();
                    reader.onloadend = () => {
                        resolve(reader.result);
                    };
                    reader.onerror = () => {
                        reject(new Error('Error reading file'));
                    };
                    reader.readAsDataURL(blob);
                });
            } else {
                base64Image = await FileSystem.readAsStringAsync(image, {
                    encoding: FileSystem.EncodingType.Base64,
                });
            }

            const token = await get(EMPLOYEE_JWT);
            await axios.post("/api/garages/logo",
                {logo: base64Image},
                {headers: {"Authorization": `Bearer ${token}`}})

            setLogoSubmitted(true);
        } catch (error) {
            setErrorMessage("Zapisanie loga nie powiodło się.")
            console.error(error);
        } finally {
            setLoading(false)
        }
    };

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

    return (
        <View className="flex-1">
            <View className="flex-row justify-between p-4 bg-gray-700">
                <Text className="text-2xl lg:text-4xl font-bold text-white lg:mt-1.5" onPress={() => {
                    router.push("/business/home")
                }}>
                    {(garage?.name ? garage.name.toUpperCase() : "")}
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

            <View className="flex-1 justify-end mb-10">
                {image && (
                    <Image
                        source={{uri: image}}
                        style={{
                            width: width > 768 ? "60%" : "80%",
                            height: width > 768 ? height * 0.5 : height * 0.4,
                            resizeMode: 'contain',
                            alignSelf: 'center',
                            marginBottom: 100
                        }}
                    />
                )}
                <CustomButton
                    title="Wybierz logo"
                    onPress={handlePickLogo}
                    containerStyles="bg-gray-700 mt-4 self-center w-3/5 lg:w-1/5"
                    textStyles="text-white font-bold"
                />
                {loading ? (
                    <ActivityIndicator size="large" color="#374151" className="mt-8"/>
                ) : (
                    <CustomButton
                        title="Zapisz logo"
                        onPress={handleSaveLogo}
                        containerStyles="bg-gray-700 mt-4 self-center w-3/5 lg:w-1/5"
                        textStyles="text-white font-bold"
                    />
                )}
                {errorMessage && (
                    <Text className="text-red-500 text-center mt-2">
                        {errorMessage}
                    </Text>
                )}
            </View>

            <Modal
                visible={logoSubmitted}
                animationType="fade"
                transparent={true}
                onRequestClose={() => setLogoSubmitted(false)}
            >
                <TouchableWithoutFeedback onPress={() => setLogoSubmitted(false)}>
                    <View className="flex-1 justify-center items-center"
                          style={{backgroundColor: 'rgba(0, 0, 0, 0.75)'}}>
                        <View className="bg-gray-700 p-5 rounded-lg w-4/5 lg:w-2/5">
                            <Text className="text-white self-center text-lg font-bold my-5">
                                Logo zostało dodane.
                            </Text>
                        </View>
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

export default LogoScreen;
