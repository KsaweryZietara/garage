import React, {useState} from "react";
import {StatusBar, Text, TextInput, View} from "react-native";
import CustomButton from "@/components/CustomButton";
import {useRouter} from "expo-router";
import axios from "axios";
import {save} from "@/utils/auth";

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

        await axios
            .post("/api/customer/login", {
                email,
                password,
            })
            .then(function (response) {
                setErrorMessage("");
                save("customer_jwt", response.data.jwt);
                router.push("/home");
            })
            .catch(function (error) {
                if (error.response.status === 400) {
                    setErrorMessage(error.response.data.message);
                } else {
                    setErrorMessage("Nieprawidłowy email lub hasło.");
                }
            });
    };

    return (
        <View className="flex-1 bg-black">
            <View className="flex-row justify-start p-4 bg-black">
                <Text className="text-white text-2xl lg:text-4xl font-bold">GARAGE</Text>
            </View>
            <View className="flex-1 justify-center items-center px-6">
                <View className="w-full max-w-xl">
                    <Text className="text-center text-3xl font-bold text-white">Logowanie</Text>
                    <Text className="text-center text-gray-400 mt-4 mb-4">
                        Jesteś tutaj nowy?{" "}
                        <Text
                            className="text-[#ff5c5c] font-bold"
                            onPress={() => router.push("/register")}
                        >
                            Zarejestruj się tutaj
                        </Text>
                    </Text>
                    <TextInput
                        placeholder="Email"
                        keyboardType="email-address"
                        value={email}
                        onChangeText={setEmail}
                        placeholderTextColor="#aaa"
                        className="bg-[#2d2d2d] text-white rounded-lg pl-4 py-3 mb-4"
                    />
                    <TextInput
                        placeholder="Hasło"
                        secureTextEntry
                        value={password}
                        onChangeText={setPassword}
                        placeholderTextColor="#aaa"
                        className="bg-[#2d2d2d] text-white rounded-lg pl-4 py-3 mb-4"
                    />
                    {errorMessage && (
                        <Text className="text-[#ff5c5c] text-center mt-2">{errorMessage}</Text>
                    )}
                    <CustomButton
                        title="Zaloguj się"
                        onPress={handleLogin}
                        containerStyles="bg-[#ff5c5c] mt-4 self-center w-3/5"
                        textStyles="text-white font-bold"
                    />
                    <Text className="text-center text-gray-400 mt-4">
                        Zapomniałeś hasła?
                    </Text>
                </View>
            </View>
            <StatusBar backgroundColor="#000000"/>
        </View>
    );
};

export default LoginScreen;
