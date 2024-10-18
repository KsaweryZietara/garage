export interface Garage {
    id: number;
    name: string;
    city: string;
    street: string;
    number: string;
    postalCode: string;
    phoneNumber: string;
    latitude: number;
    longitude: number;
    rating: number;
    distance: number;
}

export interface Employee {
    id: number;
    name: string;
    surname: string;
    confirmed: boolean;
}

export interface Service {
    id: number;
    name: string;
    time: string;
    price: string;
}

export interface Appointments {
    upcoming: CustomerAppointment[];
    inProgress: CustomerAppointment[];
    completed: CustomerAppointment[];
}

export interface CustomerAppointment {
    id: number;
    startTime: Date;
    endTime: Date;
    service: Service;
    employee: Employee;
    garage: Garage;
    rating?: number;
    comment?: string;
    car: Car;
}

export interface EmployeeAppointment {
    id: number;
    startTime: Date;
    endTime: Date;
    service: Service;
    employee?: Employee;
    car: Car;
}

export interface TimeSlot {
    startTime: Date;
    endTime: Date;
}

export interface Review {
    id: number;
    time: Date;
    service: string;
    employee: Employee;
    rating: number;
    comment?: string;
}

export interface Make {
    id: number;
    name: string;
}

export interface Model {
    id: number;
    name: string;
}

export interface Car {
    make: string;
    model: string;
}

export interface JwtPayload {
    email?: string;
    role?: string;
}
