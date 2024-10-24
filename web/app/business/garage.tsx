import {useRouter} from "expo-router";
import React, {useEffect, useState} from "react";
import {get, remove} from "@/utils/auth";
import {EMPLOYEE_JWT} from "@/constants/constants";
import axios from "axios";
import {getJwtPayload} from "@/utils/jwt";
import {
    ActivityIndicator,
    Platform, StatusBar,
    Text,
    View
} from "react-native";
import BusinessMenu from "@/components/BusinessMenu";
import {Garage} from "@/types";
import CustomButton from "@/components/CustomButton";
import CustomTextInput from "@/components/CustomTextInput";

const GarageScreen = () => {
    const router = useRouter();
    const [email, setEmail] = useState<string | null>(null);
    const [role, setRole] = useState<string | null>(null);
    const [menuVisible, setMenuVisible] = useState(false);
    const [garage, setGarage] = useState<Garage | null>(null);
    const [name, setName] = useState("");
    const [city, setCity] = useState("");
    const [street, setStreet] = useState("");
    const [number, setNumber] = useState("");
    const [postalCode, setPostalCode] = useState("");
    const [phoneNumber, setPhoneNumber] = useState("");
    const [latitude, setLatitude] = useState("");
    const [longitude, setLongitude] = useState("");
    const [loading, setLoading] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");

    useEffect(() => {
        const fetchJwtPayload = async () => {
            const jwtPayload = await getJwtPayload(EMPLOYEE_JWT);
            setEmail(jwtPayload?.email || null);
            setRole(jwtPayload?.role || null)
        };

        fetchGarage();
        fetchJwtPayload();
    }, []);

    const fetchGarage = async () => {
        const token = await get(EMPLOYEE_JWT);
        await axios.get<Garage>("/api/employees/garages", {
            headers: {"Authorization": `Bearer ${token}`}
        })
            .then((response) => {
                if (response.data) {
                    setGarage(response.data);
                    setName(response.data.name)
                    setCity(response.data.city)
                    setStreet(response.data.street)
                    setNumber(response.data.number)
                    setPostalCode(response.data.postalCode)
                    setPhoneNumber(response.data.phoneNumber)
                    setLatitude(response.data.latitude.toString())
                    setLongitude(response.data.longitude.toString())
                }
            })
            .catch((error) => {
                console.error(error)
            });
    };

    const validateGarage = () => {
        if (
            !name.trim() ||
            !city.trim() ||
            !street.trim() ||
            !number.trim() ||
            !postalCode.trim() ||
            !phoneNumber.trim() ||
            !latitude.trim() ||
            !longitude.trim()
        ) {
            return "Wszystkie pola muszą być wypełnione.";
        }

        if (name.length > 255 || city.length > 255 || street.length > 255) {
            return "Nazwa, Miasto i Ulica nie mogą przekraczać 255 znaków.";
        }

        if (number.length > 15 || postalCode.length > 15 || phoneNumber.length > 15) {
            return "Numer, kod pocztowy i numer telefonu nie mogą przekraczać 15 znaków.";
        }

        const postalCodeRegex = /^\d{2}-\d{3}$/;
        if (!postalCodeRegex.test(postalCode)) {
            return "Nieprawidłowy format kodu pocztowego.";
        }

        const phoneRegex = /^\d{9}$/;
        if (!phoneRegex.test(phoneNumber)) {
            return "Nieprawidłowy format numeru telefonu.";
        }

        const latValue = parseFloat(latitude.replace(",", "."));
        if (isNaN(latValue) || latValue < -90 || latValue > 90) {
            return "Nieprawidłowa szerokość geograficzna. Powinna być liczbą od -90 do 90.";
        }

        const lonValue = parseFloat(longitude.replace(",", "."));
        if (isNaN(lonValue) || lonValue < -180 || lonValue > 180) {
            return "Nieprawidłowa długość geograficzna. Powinna być liczbą od -180 do 180.";
        }

        return null;
    };

    const handleSubmit = async () => {
        const validationError = validateGarage();
        if (validationError) {
            setErrorMessage(validationError);
            return;
        }
        setLoading(true)
        const token = await get(EMPLOYEE_JWT);
        const latValue = parseFloat(latitude.replace(",", "."));
        const lonValue = parseFloat(longitude.replace(",", "."));
        const data = {
            name,
            city,
            street,
            number,
            postalCode,
            phoneNumber,
            latitude: latValue,
            longitude: lonValue
        };
        await axios.put("/api/garages", data, {headers: {"Authorization": `Bearer ${token}`}})
            .then(() => {
                setErrorMessage("");
                fetchGarage()
            })
            .catch((error) => {
                console.error(error)
                setErrorMessage(error.response.data.message);
            }).finally(() => {
                setLoading(false)
            });
    };

    return (
        <View className="flex-1">
            <View className="flex-row justify-between p-4 bg-gray-700">
                <Text className="text-2xl lg:text-4xl font-bold text-white lg:mt-1.5">
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

            <View className="flex-1 justify-center items-center px-6">
                <View className="w-full max-w-xl">
                    <Text className="text-3xl font-bold text-gray-700 self-center mb-8">
                        Edytuj swój garaż
                    </Text>
                    <CustomTextInput
                        placeholder="Nazwa"
                        value={name}
                        onChangeText={setName}
                    />
                    <CustomTextInput
                        placeholder="Miasto"
                        value={city}
                        onChangeText={setCity}
                    />
                    <CustomTextInput
                        placeholder="Ulica"
                        value={street}
                        onChangeText={setStreet}
                    />
                    <CustomTextInput
                        placeholder="Numer"
                        value={number}
                        onChangeText={setNumber}
                    />
                    <CustomTextInput
                        placeholder="Kod pocztowy (np. 12-345)"
                        value={postalCode}
                        onChangeText={setPostalCode}
                    />
                    <CustomTextInput
                        placeholder="Numer telefonu (np. 123456789)"
                        keyboardType="phone-pad"
                        value={phoneNumber}
                        onChangeText={setPhoneNumber}
                    />
                    <CustomTextInput
                        placeholder="Szerokość geograficzna (np. 52,237049)"
                        keyboardType="numeric"
                        value={latitude}
                        onChangeText={setLatitude}
                    />
                    <CustomTextInput
                        placeholder="Długość geograficzna (np. 21,017532)"
                        keyboardType="numeric"
                        value={longitude}
                        onChangeText={setLongitude}
                    />
                    {errorMessage && (
                        <Text className="text-red-500 text-center mt-2">
                            {errorMessage}
                        </Text>
                    )}
                    {loading ? (
                        <ActivityIndicator size="large" color="#374151"/>
                    ) : (
                        <CustomButton
                            title="Zapisz"
                            onPress={handleSubmit}
                            containerStyles="bg-gray-700 mt-4 self-center w-3/5"
                            textStyles="text-white font-bold"
                        />
                    )}
                </View>
            </View>

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

export default GarageScreen;
