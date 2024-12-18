import React from "react";
import {
    Modal,
    Platform,
    Text,
    TouchableOpacity,
    TouchableWithoutFeedback,
    View
} from "react-native";
import {useRouter} from "expo-router";
import {MECHANIC, OWNER} from "@/constants/constants";

interface MenuProps {
    menuVisible: boolean;
    onClose: () => void;
    role: string | null;
    email: string | null;
    onLogout: () => void;
}

const BusinessMenu: React.FC<MenuProps> = ({menuVisible, onClose, role, email, onLogout}) => {
    const router = useRouter();

    return (
        <Modal
            transparent={true}
            animationType="fade"
            visible={menuVisible}
            onRequestClose={onClose}
        >
            <TouchableWithoutFeedback onPress={onClose}>
                <View
                    style={{
                        flex: 1,
                        justifyContent: "flex-start",
                        alignItems: "flex-end",
                    }}
                >
                    <View
                        style={{
                            marginRight: Platform.OS === "web" ? 32 : 27,
                            marginTop: Platform.OS === "web" ? 57 : 50,
                            backgroundColor: "white",
                            borderRadius: 5,
                            padding: Platform.OS === "web" ? 12 : 6,
                            elevation: 5,
                            shadowColor: "#000",
                            shadowOffset: {width: 0, height: 2},
                            shadowOpacity: 0.25,
                            shadowRadius: 4,
                        }}
                    >
                        {role === OWNER && (
                            <View>
                                <TouchableOpacity onPress={() => router.push("/business/garage")}>
                                    <Text className="text-gray-700 font-bold mb-3">Garaż</Text>
                                </TouchableOpacity>
                                <TouchableOpacity onPress={() => router.push("/business/logo")}>
                                    <Text className="text-gray-700 font-bold mb-3">Logo</Text>
                                </TouchableOpacity>
                                <TouchableOpacity onPress={() => router.push("/business/services")}>
                                    <Text className="text-gray-700 font-bold mb-3">Usługi</Text>
                                </TouchableOpacity>
                                <TouchableOpacity onPress={() => router.push("/business/employees")}>
                                    <Text className="text-gray-700 font-bold mb-3">Pracownicy</Text>
                                </TouchableOpacity>
                            </View>
                        )}
                        {role === MECHANIC && (
                            <View>
                                <TouchableOpacity onPress={() => router.push("/business/profile-picture")}>
                                    <Text className="text-gray-700 font-bold mb-3">Zdjęcie profilowe</Text>
                                </TouchableOpacity>
                            </View>
                        )}
                        {email && (
                            <TouchableOpacity onPress={onLogout}>
                                <Text className="text-red-500 font-bold">Wyloguj</Text>
                            </TouchableOpacity>
                        )}
                    </View>
                </View>
            </TouchableWithoutFeedback>
        </Modal>
    );
};

export default BusinessMenu;
