import React, {useState} from "react";
import {View, Text, StatusBar} from "react-native";
import CustomTextInput from "@/components/CustomTextInput";
import CustomButton from "@/components/CustomButton";
import {useRouter} from "expo-router";

const NewPasswordScreen = () => {
    const router = useRouter();
    const [newPassword, setNewPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");
    const [errorMessage, setErrorMessage] = useState("");

    const validateFields = () => {
        if (!newPassword || !confirmPassword) {
            return "Wszystkie pola muszą być wypełnione.";
        }

        if (newPassword.length > 60) {
            return "Hasło nie może przekraczać 60 znaków.";
        }

        const passwordRegex = /^(?=.*[A-Z])(?=.*\d)[A-Za-z\d]{8,60}$/;
        if (!passwordRegex.test(newPassword)) {
            return "Hasło musi mieć co najmniej 8 znaków, zawierać co najmniej jedną wielką literę i jedną cyfrę.";
        }

        if (newPassword !== confirmPassword) {
            return "Hasła muszą być takie same.";
        }

        return null;
    };

    const handleNewPassword = () => {
        const validationError = validateFields();

        if (validationError) {
            setErrorMessage(validationError);
            return;
        }

        setErrorMessage("");
        router.push("/business/login");
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
                        Ustaw nowe hasło
                    </Text>
                    <Text className="text-center text-gray-500 mb-4">
                        Wprowadź swój adres email, a wyślemy Ci link do resetowania hasła.
                    </Text>
                    <CustomTextInput
                        placeholder="Nowe hasło"
                        secureTextEntry
                        value={newPassword}
                        onChangeText={setNewPassword}
                    />
                    <CustomTextInput
                        placeholder="Potwierdź nowe hasło"
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
                        title="Zmień hasło"
                        onPress={handleNewPassword}
                        containerStyles="bg-gray-700 mt-4 self-center w-3/5"
                        textStyles="text-white font-bold"
                    />
                </View>
            </View>
            <StatusBar backgroundColor="#374151"/>
        </View>
    );
};

export default NewPasswordScreen;
