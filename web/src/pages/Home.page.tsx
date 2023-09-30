import { Button, Container, Flex } from '@mantine/core';
import { HeaderSimple } from '../components/Header/HeaderSimple';
import { TableSelection } from '../components/TableLayout/TableLayout';
import { StatsGrid } from '../components/Stats/StatsGrid';
import { IAttendance, Service } from '../types/Attendance';
// import { ColorSchemeToggle } from '../components/ColorSchemeToggle/ColorSchemeToggle';

export function HomePage() {
const data: IAttendance[] = [{
    id: '1',
    author: {
      avatar: 'https://images.unsplash.com/photo-1624298357597-fd92dfbec01d?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=250&q=80',
      id: '123',
      given_name: 'Abiodun',
      family_name: 'Azeez',
      email: 'info@gmail.com',
    },
    children: 25,
    baby: 71,
    offering: 40000,
    first_timer: 20,
    teen: 50,
    service: Service.First,
    adult: 200,
    created_at: new Date().toDateString(),
  },
  {
    id: '2',
    author: {
      avatar: 'https://images.unsplash.com/photo-1632922267756-9b71242b1592?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=250&q=80',//'https://images.unsplash.com/photo-1624298357597-fd92dfbec01d?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=250&q=80',
      id: '123',
      given_name: 'Abiodun',
      family_name: 'Azeez',
      email: 'info@gmail.com',
    },
    children: 25,
    baby: 71,
    offering: 40000,
    first_timer: 20,
    teen: 50,
    service: Service.Second,
    adult: 150,
    created_at: new Date().toDateString(),
  },
];
  return (
    <>
      <HeaderSimple />
      <Container size="md" mt={2}>
        <StatsGrid />
        <Flex justify="space-between">
          <h2>Attendance</h2>
          {/* <ColorSchemeToggle /> */}
          <Button size="sm" mt={20}>Add New</Button>
        </Flex>
        <>
          {/* <TableLayout data={data} /> */}
          <TableSelection data={data} />
        </>
      </Container>
    </>
  );
}
