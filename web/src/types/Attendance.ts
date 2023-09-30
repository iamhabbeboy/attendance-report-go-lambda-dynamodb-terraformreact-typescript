import { IUser } from './User';

export interface IAttendance {
    id: string;
    author: IUser,
    first_timer: number;
    offering: number;
    service: string;
    children: number;
    baby: number;
    adult: number;
    teen: number;
    categories?: Category[];
    created_at: string;
    updated_at?: string;
    delete_at?: string;
}

export interface Category {
    [key: string]: string;
}

export enum Service {
    First = 'First',
    Second = 'Second',
    Third = 'Third',
}
