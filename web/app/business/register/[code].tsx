import React, {useState} from "react";
import {View, Text, StatusBar} from "react-native";
import CustomTextInput from "@/components/CustomTextInput";
import CustomButton from "@/components/CustomButton";
import {useLocalSearchParams, useRouter} from "expo-router";
import axios from "axios";

const CreateMechanicAccountScreen = () => {
    const {code} = useLocalSearchParams()
    const router = useRouter();
    const [name, setName] = useState("");
    const [surname, setSurname] = useState("");
    const [password, setPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");
    const [errorMessage, setErrorMessage] = useState("");

    const validateFields = () => {
        if (
            !name.trim() ||
            !surname.trim() ||
            !password ||
            !confirmPassword
        ) {
            return "Wszystkie pola muszą być wypełnione.";
        }

        if (name.length > 255 || surname.length > 255 || password.length > 255 || confirmPassword.length > 255) {
            return "Imię, Nazwisko i Email nie mogą przekraczać 40 znaków.";
        }

        if (password.length > 60) {
            return "Hasło nie może przekraczać 60 znaków.";
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

    const handleCreateAccount = async () => {
        const validationError = validateFields();

        if (validationError) {
            setErrorMessage(validationError);
            return;
        }

        await axios.post(`/api/employees/register/${code}`, {
            name,
            surname,
            password,
            confirmPassword
        })
            .then(() => {
                setErrorMessage("");
                router.push("/business/login");
            })
            .catch((error) => {
                console.error(error)
                setErrorMessage(error.response.data.message);
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
                    <Text className="text-center text-xl font-bold mb-6 text-gray-700">
                        Utwórz konto
                    </Text>
                    <CustomTextInput
                        placeholder="Imię"
                        value={name}
                        onChangeText={setName}
                    />
                    <CustomTextInput
                        placeholder="Nazwisko"
                        value={surname}
                        onChangeText={setSurname}
                    />
                    <CustomTextInput
                        placeholder="Hasło"
                        secureTextEntry
                        value={password}
                        onChangeText={setPassword}
                    />
                    <CustomTextInput
                        placeholder="Potwierdź hasło"
                        secureTextEntry
                        value={confirmPassword}
                        onChangeText={setConfirmPassword}
                    />
                    {errorMessage && (
                        <Text className="text-red-500 text-center mt-2">
                            {errorMessage}
                        </Text>
                    )}
                    <CustomButton
                        title="Utwórz konto"
                        onPress={handleCreateAccount}
                        containerStyles="bg-gray-700 mt-4 self-center w-3/5"
                        textStyles="text-white font-bold"
                    />
                    <Text className="text-center text-gray-500 mt-4">
                        Masz już konto?{" "}
                        <Text
                            className="text-gray-700"
                            onPress={() => router.push("/business/login")}
                        >
                            Zaloguj się tutaj
                        </Text>
                    </Text>
                </View>
            </View>
            <StatusBar backgroundColor="#374151"/>
        </View>
    );
};

export default CreateMechanicAccountScreen;
