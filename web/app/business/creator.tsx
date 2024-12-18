import React, {useEffect, useState} from "react";
import {
    View,
    Text,
    StatusBar,
    ScrollView,
    Keyboard, ActivityIndicator
} from "react-native";
import CustomTextInput from "@/components/CustomTextInput";
import CustomButton from "@/components/CustomButton";
import uuid from "react-native-uuid";
// @ts-ignore
import {ProgressSteps, ProgressStep} from "react-native-progress-steps";
import axios from "axios";
import {get} from "@/utils/auth";
import {useRouter} from "expo-router";
import {EMPLOYEE_JWT} from "@/constants/constants";

interface Service {
    id: string;
    name: string;
    time: string;
    price: string;
}

const CreatorScreen = () => {
    const router = useRouter();
    const [name, setName] = useState("");
    const [city, setCity] = useState("");
    const [street, setStreet] = useState("");
    const [number, setNumber] = useState("");
    const [postalCode, setPostalCode] = useState("");
    const [phoneNumber, setPhoneNumber] = useState("");
    const [latitude, setLatitude] = useState("");
    const [longitude, setLongitude] = useState("");

    const [services, setServices] = useState<Service[]>([]);
    const [serviceName, setServiceName] = useState("");
    const [serviceTime, setServiceTime] = useState("");
    const [servicePrice, setServicePrice] = useState("");

    const [employeeEmail, setEmployeeEmail] = useState("");
    const [employeeEmails, setEmployeeEmails] = useState<string[]>([]);

    const [errorMessage, setErrorMessage] = useState("");
    const [errors, setErrors] = useState(true);

    const [loading, setLoading] = useState(false);

    const [buttonsVisible, setButtonsVisible] = useState(false)
    useEffect(() => {
        const showSubscription = Keyboard.addListener("keyboardDidShow", () => {
            setButtonsVisible(true);
        })
        const hideSubscription = Keyboard.addListener("keyboardDidHide", () => {
            setButtonsVisible(false)
        })

        return () => {
            showSubscription.remove();
            hideSubscription.remove();
        }
    }, [])

    const validateGarage = () => {
        if (
            !name.trim() ||
            !city.trim() ||
            !street.trim() ||
            !number.trim() ||
            !postalCode.trim() ||
            !phoneNumber.trim() ||
            !latitude.trim() ||
            !longitude.trim()
        ) {
            return "Wszystkie pola muszą być wypełnione.";
        }

        if (name.length > 255 || city.length > 255 || street.length > 255) {
            return "Nazwa, Miasto i Ulica nie mogą przekraczać 255 znaków.";
        }

        if (number.length > 15 || postalCode.length > 15 || phoneNumber.length > 15) {
            return "Numer, kod pocztowy i numer telefonu nie mogą przekraczać 15 znaków.";
        }

        const postalCodeRegex = /^\d{2}-\d{3}$/;
        if (!postalCodeRegex.test(postalCode)) {
            return "Nieprawidłowy format kodu pocztowego.";
        }

        const phoneRegex = /^\d{9}$/;
        if (!phoneRegex.test(phoneNumber)) {
            return "Nieprawidłowy format numeru telefonu.";
        }

        const latValue = parseFloat(latitude.replace(",", "."));
        if (isNaN(latValue) || latValue < -90 || latValue > 90) {
            return "Nieprawidłowa szerokość geograficzna. Powinna być liczbą od -90 do 90.";
        }

        const lonValue = parseFloat(longitude.replace(",", "."));
        if (isNaN(lonValue) || lonValue < -180 || lonValue > 180) {
            return "Nieprawidłowa długość geograficzna. Powinna być liczbą od -180 do 180.";
        }

        return null;
    };

    const handleGarage = () => {
        const validationError = validateGarage();

        if (validationError) {
            setErrors(true);
            setErrorMessage(validationError);
            return;
        }

        setErrors(false);
        setErrorMessage("");
    };

    const addService = () => {
        if (!serviceName || !serviceTime || !servicePrice) {
            setErrorMessage("Wszystkie pola dla usługi muszą być wypełnione.");
            return;
        }

        if (serviceName.length > 255) {
            setErrorMessage("Nazwa nie może przekraczać 255 znaków.");
            return;
        }

        const timeNumber = Number(serviceTime);
        const priceNumber = Number(servicePrice);

        if (isNaN(timeNumber) || timeNumber <= 0 || !Number.isInteger(timeNumber)) {
            setErrorMessage("Czas musi być podany w pełnych godzinach.");
            return;
        }

        if (timeNumber > 720) {
            setErrorMessage("Czas usługi nie może być dłuższy niż miesiąc.");
            return;
        }

        if (isNaN(priceNumber) || priceNumber <= 0 || !Number.isInteger(priceNumber)) {
            setErrorMessage("Cena musi być podana w pełnych złotówkach.");
            return;
        }

        const newService = {
            id: uuid.v4().toString(),
            name: serviceName,
            time: serviceTime,
            price: servicePrice,
        };

        setServices([...services, newService]);
        setServiceName("");
        setServiceTime("");
        setServicePrice("");
        setErrorMessage("");
    };

    const removeService = (id: string) => {
        setServices(services.filter((service) => service.id !== id));
    };

    const addEmployeeEmail = () => {
        if (!employeeEmail.trim()) {
            setErrorMessage("Pole email nie może być puste.");
            return;
        }

        if (employeeEmail.length > 255) {
            return "Email nie może przekraczać 255 znaków.";
        }

        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        if (!emailRegex.test(employeeEmail)) {
            setErrorMessage("Nieprawidłowy format adresu e-mail.");
            return;
        }

        if (employeeEmails.includes(employeeEmail)) {
            setErrorMessage("Email już istnieje.");
            return;
        }

        setEmployeeEmails([...employeeEmails, employeeEmail]);
        setEmployeeEmail("");
        setErrorMessage("");
    };

    const removeEmployeeEmail = (email: string) => {
        setEmployeeEmails(employeeEmails.filter((e) => e !== email));
    };

    const handleSubmit = async () => {
        const token = await get(EMPLOYEE_JWT);
        const latValue = parseFloat(latitude.replace(",", "."));
        const lonValue = parseFloat(longitude.replace(",", "."));
        const data = {
            name,
            city,
            street,
            number,
            postalCode,
            phoneNumber,
            latitude: latValue,
            longitude: lonValue,
            services: services.map(service => ({
                name: service.name,
                time: parseInt(service.time, 10),
                price: parseInt(service.price, 10)
            })),
            employeeEmails
        };
        setLoading(true)
        await axios.post("/api/garages", data, {headers: {"Authorization": `Bearer ${token}`}})
            .then(() => {
                setErrorMessage("");
                router.push("/business/home")
            })
            .catch((error) => {
                console.error(error)
                setErrorMessage(error.response.data.message);
            }).finally(() => {
                setLoading(false)
            });
    };

    return (
        <View className="flex-1 bg-white">
            <View className="flex-row justify-start p-4">
                <Text className="text-2xl lg:text-4xl font-bold text-gray-700 lg:mt-1.5">
                    GARAGE DLA BIZNESU
                </Text>
            </View>
            <View style={{flex: 1}}>
                {loading ? (
                    <View className="flex-1 justify-center items-center">
                        <Text className="text-gray-700 text-xl mb-6">Proszę czekać, tworzymy twój warsztat...</Text>
                        <ActivityIndicator size="large" color="#374151" className=""/>
                    </View>
                ) : (
                    <ProgressSteps
                        activeStepIconBorderColor="#374151"
                        completedProgressBarColor="#374151"
                        completedStepIconColor="#374151"
                        activeLabelColor="#374151"
                    >
                        <ProgressStep
                            label="Dane warsztatu"
                            nextBtnText="Dalej"
                            nextBtnTextStyle={{
                                color: "#FFFFFF",
                                fontSize: 18,
                                backgroundColor: "#374151",
                                borderRadius: 12,
                                padding: 12,
                                minWidth: 100,
                                textAlign: "center",
                            }}
                            onNext={handleGarage}
                            errors={errors}
                            removeBtnRow={buttonsVisible}
                        >
                            <View className="flex-1 justify-center items-center px-6">
                                <View className="w-full max-w-xl">
                                    <CustomTextInput
                                        placeholder="Nazwa"
                                        value={name}
                                        onChangeText={setName}
                                    />
                                    <CustomTextInput
                                        placeholder="Miasto"
                                        value={city}
                                        onChangeText={setCity}
                                    />
                                    <CustomTextInput
                                        placeholder="Ulica"
                                        value={street}
                                        onChangeText={setStreet}
                                    />
                                    <CustomTextInput
                                        placeholder="Numer"
                                        value={number}
                                        onChangeText={setNumber}
                                    />
                                    <CustomTextInput
                                        placeholder="Kod pocztowy (np. 12-345)"
                                        value={postalCode}
                                        onChangeText={setPostalCode}
                                    />
                                    <CustomTextInput
                                        placeholder="Numer telefonu (np. 123456789)"
                                        keyboardType="phone-pad"
                                        value={phoneNumber}
                                        onChangeText={setPhoneNumber}
                                    />
                                    <CustomTextInput
                                        placeholder="Szerokość geograficzna (np. 52,237049)"
                                        keyboardType="numeric"
                                        value={latitude}
                                        onChangeText={setLatitude}
                                    />
                                    <CustomTextInput
                                        placeholder="Długość geograficzna (np. 21,017532)"
                                        keyboardType="numeric"
                                        value={longitude}
                                        onChangeText={setLongitude}
                                    />
                                    {errorMessage && (
                                        <Text className="text-red-500 text-center mt-2">
                                            {errorMessage}
                                        </Text>
                                    )}
                                </View>
                            </View>
                        </ProgressStep>
                        <ProgressStep
                            label="Usługi"
                            previousBtnText="Cofnij"
                            nextBtnText="Dalej"
                            nextBtnTextStyle={{
                                color: "#FFFFFF",
                                fontSize: 18,
                                backgroundColor: "#374151",
                                borderRadius: 12,
                                padding: 12,
                                minWidth: 100,
                                textAlign: "center",
                                marginLeft: 20,
                            }}
                            previousBtnTextStyle={{color: "#374151", padding: 12}}
                            removeBtnRow={buttonsVisible}
                        >
                            <View className="flex-1 justify-center items-center px-6 mt-6">
                                <View className="w-full max-w-xl">
                                    <CustomTextInput
                                        placeholder="Nazwa usługi"
                                        value={serviceName}
                                        onChangeText={setServiceName}
                                    />
                                    <CustomTextInput
                                        placeholder="Czas wykonania (w godzinach)"
                                        value={serviceTime}
                                        onChangeText={setServiceTime}
                                        keyboardType="numeric"
                                    />
                                    <CustomTextInput
                                        placeholder="Cena (zł)"
                                        value={servicePrice}
                                        onChangeText={setServicePrice}
                                        keyboardType="numeric"
                                    />
                                    <CustomButton
                                        title="Dodaj usługę"
                                        onPress={addService}
                                        containerStyles="bg-gray-700 mt-4 mb-4 self-center w-3/5"
                                        textStyles="text-white font-bold"
                                    />
                                    {errorMessage && (
                                        <Text className="text-red-500 text-center mt-2">
                                            {errorMessage}
                                        </Text>
                                    )}
                                    <ScrollView>
                                        {services.map((item) => (
                                            <View
                                                key={item.id.toString()}
                                                className="flex-row justify-between items-center bg-gray-100 p-4 my-2 rounded-lg shadow-sm">
                                                <View>
                                                    <Text className="text-lg font-semibold text-gray-800">
                                                        {item.name}
                                                    </Text>
                                                    <Text className="text-sm text-gray-500">
                                                        {item.time} godz. - {item.price} zł
                                                    </Text>
                                                </View>
                                                <Text className="text-red-600 font-bold"
                                                      onPress={() => removeService(item.id)}>Usuń</Text>
                                            </View>
                                        ))}
                                    </ScrollView>
                                </View>
                            </View>
                        </ProgressStep>
                        <ProgressStep
                            label="Mechanicy"
                            previousBtnText="Cofnij"
                            finishBtnText="Utwórz warsztat"
                            nextBtnTextStyle={{
                                color: "#FFFFFF",
                                fontSize: 18,
                                backgroundColor: "#374151",
                                borderRadius: 12,
                                padding: 12,
                                minWidth: 100,
                                textAlign: "center",
                            }}
                            previousBtnTextStyle={{color: "#374151", padding: 12}}
                            onSubmit={handleSubmit}
                            removeBtnRow={buttonsVisible}
                        >
                            <View className="flex-1 justify-center items-center px-6 mt-6">
                                <View className="w-full max-w-xl">
                                    <CustomTextInput
                                        placeholder="Email pracownika"
                                        value={employeeEmail}
                                        onChangeText={setEmployeeEmail}
                                        keyboardType="email-address"
                                    />
                                    <CustomButton
                                        title="Dodaj email"
                                        onPress={addEmployeeEmail}
                                        containerStyles="bg-gray-700 mt-4 mb-4 self-center w-3/5"
                                        textStyles="text-white font-bold"
                                    />
                                    {errorMessage && (
                                        <Text className="text-red-500 text-center mt-2">
                                            {errorMessage}
                                        </Text>
                                    )}
                                    <ScrollView>
                                        {employeeEmails.map((email, index) => (
                                            <View
                                                key={index.toString()}
                                                className="flex-row justify-between items-center bg-gray-100 p-4 my-2 rounded-lg shadow-sm">
                                                <Text className="text-lg font-semibold text-gray-800">
                                                    {email}
                                                </Text>
                                                <Text className="text-red-600 font-bold"
                                                      onPress={() => removeEmployeeEmail(email)}>Usuń</Text>
                                            </View>
                                        ))}
                                    </ScrollView>
                                </View>
                            </View>
                        </ProgressStep>
                    </ProgressSteps>
                )}
            </View>
            <StatusBar backgroundColor="#374151"/>
        </View>
    );
};

export default CreatorScreen;
