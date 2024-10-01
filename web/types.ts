export interface Garage {
    id: number;
    name: string;
    city: string;
    street: string;
    number: string;
    postalCode: string;
    phoneNumber: string;
}

export interface Employee {
    id: number;
    name: string;
    surname: string;
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
}

export interface EmployeeAppointment {
    id: number;
    startTime: Date;
    endTime: Date;
    service: Service;
    employee?: Employee;
}

export interface TimeSlot {
    startTime: Date;
    endTime: Date;
}