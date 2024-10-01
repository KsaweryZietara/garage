import React, {useState} from "react";
import {View, Text, StatusBar, TextInput} from "react-native";
import CustomButton from "@/components/CustomButton";
import {useRouter} from "expo-router";
import axios from "axios";

const RegisterScreen = () => {
    const router = useRouter();
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");
    const [errorMessage, setErrorMessage] = useState("");

    const validateFields = () => {
        if (
            !email.trim() ||
            !password ||
            !confirmPassword
        ) {
            return "Wszystkie pola muszą być wypełnione.";
        }

        if (email.length > 255 || password.length > 255 || confirmPassword.length > 255) {
            return "Imię, Nazwisko i Email nie mogą przekraczać 40 znaków.";
        }

        if (password.length > 60) {
            return "Hasło nie może przekraczać 60 znaków.";
        }

        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        if (!emailRegex.test(email)) {
            return "Nieprawidłowy format adresu e-mail.";
        }

        const passwordRegex = /^(?=.*[A-Z])(?=.*\d)[A-Za-z\d]{8,60}$/;
        if (!passwordRegex.test(password)) {
            return "Hasło musi mieć co najmniej 8 znaków, zawierać co najmniej jedną wielką literę i jedną cyfrę.";
        }

        if (password !== confirmPassword) {
            return "Hasła muszą być takie same.";
        }

        return null;
    };

    const handleRegister = async () => {
        const validationError = validateFields();

        if (validationError) {
            setErrorMessage(validationError);
            return;
        }

        await axios.post("/api/customers/register", {
            email,
            password,
            confirmPassword
        })
            .then(() => {
                setErrorMessage("");
                router.push("/login");
            })
            .catch((error) => {
                console.error(error)
                setErrorMessage(error.response.data.message);
            });
    };

    return (
        <View className="flex-1 bg-black">
            <View className="flex-row justify-start p-4 bg-black">
                <Text className="text-white text-2xl lg:text-4xl font-bold">GARAGE</Text>
            </View>
            <View className="flex-1 justify-center items-center px-6">
                <View className="w-full max-w-xl">
                    <Text className="text-center text-3xl font-bold text-white mb-4">Utwórz konto</Text>
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
                    <TextInput
                        placeholder="Potwierdź hasło"
                        secureTextEntry
                        value={confirmPassword}
                        onChangeText={setConfirmPassword}
                        placeholderTextColor="#aaa"
                        className="bg-[#2d2d2d] text-white rounded-lg pl-4 py-3 mb-4"
                    />
                    {errorMessage && (
                        <Text className="text-red-500 text-center mt-2">{errorMessage}</Text>
                    )}
                    <CustomButton
                        title="Utwórz konto"
                        onPress={handleRegister}
                        containerStyles="bg-red-500 mt-4 self-center w-3/5"
                        textStyles="text-white font-bold"
                    />
                    <Text className="text-center text-gray-400 mt-4">
                        Masz już konto?{" "}
                        <Text
                            className="text-red-500 font-bold"
                            onPress={() => router.push("/login")}
                        >
                            Zaloguj się tutaj
                        </Text>
                    </Text>
                </View>
            </View>
            <StatusBar backgroundColor="#000000"/>
        </View>
    );
};

export default RegisterScreen;
