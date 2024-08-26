import React, {useState} from "react";
import {
    View,
    Text,
    StatusBar,
    ScrollView
} from "react-native";
import CustomTextInput from "@/components/CustomTextInput";
import CustomButton from "@/components/CustomButton";
import uuid from "react-native-uuid";
import CheckBox from "react-native-check-box"
// @ts-ignore
import {ProgressSteps, ProgressStep} from "react-native-progress-steps";

interface Service {
    id: string;
    name: string;
    time: string;
    price: string;
}

const CreatorScreen = () => {
    const [name, setName] = useState("");
    const [city, setCity] = useState("");
    const [street, setStreet] = useState("");
    const [number, setNumber] = useState("");
    const [postalCode, setPostalCode] = useState("");
    const [phoneNumber, setPhoneNumber] = useState("");

    const [services, setServices] = useState<Service[]>([]);
    const [serviceName, setServiceName] = useState("");
    const [serviceTime, setServiceTime] = useState("");
    const [servicePrice, setServicePrice] = useState("");

    const [isMechanic, setIsMechanic] = useState(false);
    const [employeeEmail, setEmployeeEmail] = useState("");
    const [employeeEmails, setEmployeeEmails] = useState<string[]>([]);

    const [errorMessage, setErrorMessage] = useState("");
    const [errors, setErrors] = useState(true);

    const validateGarage = () => {
        if (
            !name.trim() ||
            !city.trim() ||
            !street.trim() ||
            !number.trim() ||
            !postalCode.trim() ||
            !phoneNumber.trim()
        ) {
            return "Wszystkie pola muszą być wypełnione.";
        }

        if (name.length > 50 || city.length > 50 || street.length > 50) {
            return "Nazwa, Miasto i Ulica nie mogą przekraczać 50 znaków.";
        }

        const postalCodeRegex = /^\d{2}-\d{3}$/;
        if (!postalCodeRegex.test(postalCode)) {
            return "Nieprawidłowy format kodu pocztowego.";
        }

        const phoneRegex = /^\+?\d{9,15}$/;
        if (!phoneRegex.test(phoneNumber)) {
            return "Nieprawidłowy format numeru telefonu.";
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

        if (serviceName.length > 200) {
            return "Nazwa nie mogą przekraczać 200 znaków.";
        }

        const newService = {
            id: uuid.v4().toString(),
            name: serviceName,
            time: serviceTime,
            price: servicePrice,
        };

        setServices([...services, newService]);
        setServiceName("")
        setServiceTime("")
        setServicePrice("")
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

    const handleSubmit = () => {
        console.log("submit");
    };

    return (
        <View className="flex-1 bg-white">
            <View className="flex-row justify-start p-4">
                <Text className="text-lg lg:text-4xl font-bold text-gray-700">
                    GARAGE DLA BIZNESU
                </Text>
            </View>
            <View style={{flex: 1}}>
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
                    >
                        <View className="flex-1 justify-center items-center px-6 mt-12">
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
                                    placeholder="Kod pocztowy"
                                    value={postalCode}
                                    onChangeText={setPostalCode}
                                />
                                <CustomTextInput
                                    placeholder="Numer telefonu"
                                    keyboardType="phone-pad"
                                    value={phoneNumber}
                                    onChangeText={setPhoneNumber}
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
                    >
                        <View className="flex-1 justify-center items-center px-6 mt-6">
                            <View className="w-full max-w-xl">
                                <CustomTextInput
                                    placeholder="Nazwa usługi"
                                    value={serviceName}
                                    onChangeText={setServiceName}
                                />
                                <CustomTextInput
                                    placeholder="Czas wykonania (min)"
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
                                                    {item.time} min - {item.price} zł
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
                    >
                        <View className="flex-1 justify-center items-center px-6 mt-6">
                            <View className="w-full max-w-xl">
                                <CustomTextInput
                                    placeholder="Email pracownika"
                                    value={employeeEmail}
                                    onChangeText={setEmployeeEmail}
                                    keyboardType="email-address"
                                />
                                <CheckBox
                                    style={{flex: 1, marginBottom: 10, marginTop: 10}}
                                    onClick={() => {
                                        setIsMechanic(!isMechanic)
                                    }}
                                    isChecked={isMechanic}
                                    rightText="Jestem mechanikiem"
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
            </View>
            <StatusBar backgroundColor="#374151"/>
        </View>
    );
};

export default CreatorScreen;
