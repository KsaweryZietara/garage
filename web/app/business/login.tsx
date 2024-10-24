import React, {useState} from "react";
import {ActivityIndicator, StatusBar, Text, View} from "react-native";
import CustomTextInput from "@/components/CustomTextInput";
import CustomButton from "@/components/CustomButton";
import {useRouter} from "expo-router";
import axios from "axios";
import {save} from "@/utils/auth";
import {EMPLOYEE_JWT} from "@/constants/constants";

const LoginScreen = () => {
    const router = useRouter();
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [errorMessage, setErrorMessage] = useState("");
    const [loading, setLoading] = useState(false);

    const validateFields = () => {
        if (!email.trim() || !password) {
            return "Wszystkie pola muszą być wypełnione.";
        }

        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        if (!emailRegex.test(email)) {
            return "Nieprawidłowy format adresu e-mail.";
        }

        return null;
    };

    const handleLogin = async () => {
        const validationError = validateFields();

        if (validationError) {
            setErrorMessage(validationError);
            return;
        }

        setLoading(true)

        await axios.post("/api/employees/login", {
            email,
            password,
        })
            .then((response) => {
                setErrorMessage("");
                save(EMPLOYEE_JWT, response.data.jwt)

                axios.get("/api/employees/garages", {headers: {"Authorization": `Bearer ${response.data.jwt}`}})
                    .then(() => {
                        router.push("/business/home")
                    })
                    .catch((error) => {
                        console.error(error)
                        router.push("/business/creator")
                    })

            })
            .catch((error) => {
                console.error(error)
                if (error.response.status === 400) {
                    setErrorMessage(error.response.data.message);
                } else {
                    setErrorMessage("Nieprawidłowy email lub hasło.")
                }
            }).finally(() => {
                setLoading(false)
            });
    };

    return (
        <View className="flex-1 bg-white">
            <View className="flex-row justify-start p-4">
                <Text className="text-2xl lg:text-4xl font-bold text-gray-700 lg:mt-1.5">
                    GARAGE DLA BIZNESU
                </Text>
            </View>
            <View className="flex-1 justify-center items-center px-6">
                <View className="w-full max-w-xl">
                    <Text className="text-center text-xl font-bold text-gray-700">
                        Logowanie
                    </Text>
                    <Text className="text-center text-gray-500 mt-4 mb-4">
                        Jesteś tutaj nowy?{" "}
                        <Text
                            className="text-gray-700"
                            onPress={() => router.push("/business/register")}
                        >
                            Zarejestruj się tutaj
                        </Text>
                    </Text>
                    <CustomTextInput
                        placeholder="Email"
                        keyboardType="email-address"
                        value={email}
                        onChangeText={setEmail}
                    />
                    <CustomTextInput
                        placeholder="Hasło"
                        secureTextEntry
                        value={password}
                        onChangeText={setPassword}
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
                            title="Zaloguj się"
                            onPress={handleLogin}
                            containerStyles="bg-gray-700 mt-4 self-center w-3/5"
                            textStyles="text-white font-bold"
                        />
                    )}
                </View>
            </View>
            <StatusBar backgroundColor="#374151"/>
        </View>
    );
};

export default LoginScreen;
