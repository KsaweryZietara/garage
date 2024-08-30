import React, {useState} from "react";
import {Platform, StatusBar, Text, View} from "react-native";
import CustomTextInput from "@/components/CustomTextInput";
import CustomButton from "@/components/CustomButton";
import {useRouter} from "expo-router";
import axios from "axios";
import * as SecureStore from 'expo-secure-store';
import Cookies from 'js-cookie';

async function save(key: string, value: string) {
    if (Platform.OS === "android" || Platform.OS === "ios") {
        await SecureStore.setItemAsync(key, value);
    } else {
        Cookies.set(key, value);
    }
}

async function getValueFor(key: string) {
    if (Platform.OS === "android" || Platform.OS === "ios") {
        return await SecureStore.getItemAsync(key)
    } else {
        return Cookies.get(key)
    }
}

const LoginScreen = () => {
    const router = useRouter();
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [errorMessage, setErrorMessage] = useState("");

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

        await axios.post('/api/business/login', {
            email,
            password,
        })
            .then(function (response) {
                setErrorMessage("");
                save("jwt", response.data.jwt)
                router.push("/business/creator");
            })
            .catch(function (error) {
                if (error.response.status === 400) {
                    setErrorMessage(error.response.data.message);
                } else {
                    setErrorMessage("Nieprawidłowy email lub hasło.")
                }
            });
    };

    return (
        <View className="flex-1 bg-white">
            <View className="flex-row justify-start p-4">
                <Text className="text-lg lg:text-4xl font-bold text-gray-700">
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
                    <CustomButton
                        title="Zaloguj się"
                        onPress={handleLogin}
                        containerStyles="bg-gray-700 mt-4 self-center w-3/5"
                        textStyles="text-white font-bold"
                    />
                    <Text className="text-center text-gray-700 mt-4"
                          onPress={() => router.push("/business/recover-password")}>
                        Zapomniałeś hasła?
                    </Text>
                </View>
            </View>
            <StatusBar backgroundColor="#374151"/>
        </View>
    );
};

export default LoginScreen;
