import React from "react";
import {TextInput, TextInputProps} from "react-native";

interface CustomTextInputProps extends TextInputProps {
    inputStyles?: string;
}

const CustomTextInput = ({
                             inputStyles = "",
                             ...props
                         }: CustomTextInputProps) => {
    return (
        <TextInput
            className={`border-b border-gray-400 py-2 mb-4 ${inputStyles}`}
            {...props}
        />
    );
};

export default CustomTextInput;
