import React, {useState} from "react";
import {View, Text, StatusBar} from "react-native";
import CustomTextInput from "@/components/CustomTextInput";
import CustomButton from "@/components/CustomButton";

const RecoverPasswordScreen = () => {
    const [email, setEmail] = useState("");
    const [errorMessage, setErrorMessage] = useState("");
    const [emailSent, setEmailSent] = useState(false);

    const validateFields = () => {
        if (!email.trim()) {
            return "Email nie może być pusty.";
        }

        return null;
    };

    const handleRecoverPassword = () => {
        const validationError = validateFields();

        if (validationError) {
            setErrorMessage(validationError);
            return;
        }

        setEmailSent(true);
        setErrorMessage("");
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
                    {emailSent ? (
                        <Text className="text-center text-xl font-bold text-gray-700">
                            Link został wysłany na podany adres email.
                        </Text>
                    ) : (
                        <>
                            <Text className="text-center text-xl font-bold text-gray-700">
                                Zapomniałeś hasła?
                            </Text>
                            <Text className="text-center text-gray-500 mt-4 mb-4">
                                Wprowadź swój adres email, a wyślemy Ci link do resetowania
                                hasła.
                            </Text>
                            <CustomTextInput
                                placeholder="Email"
                                keyboardType="email-address"
                                value={email}
                                onChangeText={setEmail}
                            />
                            {errorMessage && (
                                <Text className="text-red-500 text-center mt-2">
                                    {errorMessage}
                                </Text>
                            )}
                            <CustomButton
                                title="Przypomnij"
                                onPress={handleRecoverPassword}
                                containerStyles="bg-gray-700 mt-4 self-center w-3/5"
                                textStyles="text-white font-bold"
                            />
                        </>
                    )}
                </View>
            </View>
            <StatusBar backgroundColor="#374151"/>
        </View>
    );
};

export default RecoverPasswordScreen;
