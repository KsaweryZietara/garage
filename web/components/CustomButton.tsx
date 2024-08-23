import React from "react";
import {Text, TouchableOpacity} from "react-native";

interface CustomButtonProps {
    onPress: () => void;
    title: string;
    textStyles?: string;
    containerStyles?: string;
}

const CustomButton = ({
                          onPress,
                          title,
                          textStyles = "",
                          containerStyles = "",
                      }: CustomButtonProps) => {
    return (
        <TouchableOpacity
            activeOpacity={0.7}
            className={`rounded-xl bg-gray-400 py-3 justify-center items-center ${containerStyles} `}
            onPress={onPress}
        >
            <Text
                className={`text-primary text-lg ${textStyles}`}
            >
                {title}
            </Text>
        </TouchableOpacity>
    );
};

export default CustomButton;
