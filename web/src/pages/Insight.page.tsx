import { Button, Container } from '@mantine/core';
import { IconChevronLeft } from '@tabler/icons-react';
import { HeaderSimple } from '../components/Header/HeaderSimple';

export function InsightPage() {
  return (
    <>
      <HeaderSimple />
      <Container size="md" mt={2}>
      <Button variant="light"><IconChevronLeft /></Button>
        <h1 style={{ color: '#666' }}>First Service - 2023-02-01 </h1>
      </Container>
    </>
 );
}
export default InsightPage;
