import React from 'react';
import {Modal, TouchableOpacity, TouchableWithoutFeedback, View, Text, Platform} from 'react-native';
import {remove} from '@/utils/auth';
import {useRouter} from "expo-router";

interface MenuModalProps {
    visible: boolean;
    onClose: () => void;
    email: string | null;
    setEmail: (email: string | null) => void;
}

const MenuModal: React.FC<MenuModalProps> = ({visible, onClose, email, setEmail}) => {
    const router = useRouter()

    const handleLogout = () => {
        remove("customer_jwt");
        setEmail(null);
        onClose();
        router.push("/home")
    };

    return (
        <Modal
            transparent={true}
            animationType="fade"
            visible={visible}
            onRequestClose={onClose}
        >
            <TouchableWithoutFeedback onPress={onClose}>
                <View style={{
                    flex: 1,
                    justifyContent: 'flex-start',
                    alignItems: 'flex-end',
                }}>
                    <View style={{
                        marginRight: Platform.OS === 'web' ? 32 : 27,
                        marginTop: Platform.OS === 'web' ? 52 : 50,
                        backgroundColor: 'white',
                        borderRadius: 5,
                        padding: Platform.OS === 'web' ? 12 : 6,
                        elevation: 5,
                        shadowColor: '#000',
                        shadowOffset: {
                            width: 0,
                            height: 2,
                        },
                        shadowOpacity: 0.25,
                        shadowRadius: 4,
                    }}>
                        {email && (
                            <View>
                                <TouchableOpacity onPress={() => router.push("/appointments")}>
                                    <Text className="text-black font-bold">Moje wizyty</Text>
                                </TouchableOpacity>
                                <TouchableOpacity onPress={handleLogout}>
                                    <Text className="text-red-500 font-bold mt-3">Wyloguj</Text>
                                </TouchableOpacity>
                            </View>
                        )}
                    </View>
                </View>
            </TouchableWithoutFeedback>
        </Modal>
    );
};

export default MenuModal;
