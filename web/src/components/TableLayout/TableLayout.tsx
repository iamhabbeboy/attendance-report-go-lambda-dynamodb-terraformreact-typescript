import cx from 'clsx';
import { useState } from 'react';
import { Table, Checkbox, ScrollArea, Group, Avatar, Text, rem, Button } from '@mantine/core';
import classes from './tablelayout.module.css';
import { IAttendance } from '../../types/Attendance';
import { IconChevronDown, IconChevronRight } from '@tabler/icons-react';
import { useNavigate } from 'react-router-dom';


interface TableSelectionProps {
  data: IAttendance[]
}
export function TableSelection({ data }: TableSelectionProps) {
  const navigate = useNavigate();
  const [selection, setSelection] = useState(['1']);
  const toggleRow = (id: string) =>
    setSelection((current) =>
      current.includes(id) ? current.filter((item) => item !== id) : [...current, id]
    );
  const toggleAll = () =>
    setSelection((current) => (current.length === data.length ? [] : data.map((item) => item.id)));

  const rows = data.map((item) => {
    const selected = selection.includes(item.id);

  const handleInsight = () => navigate('/insight/1');

    return (
      <Table.Tr key={item.id} className={cx({ [classes.rowSelected]: selected })}>
        <Table.Td>
          <Checkbox checked={selection.includes(item.id)} onChange={() => toggleRow(item.id)} />
        </Table.Td>
        <Table.Td>
          <Group gap="sm">
            <Avatar size={26} src={item.author.avatar} radius={26} />
            <Text size="sm" fw={500}>
              Abiodun A.
            </Text>
          </Group>
        </Table.Td>
        <Table.Td>{item.service}</Table.Td>
        <Table.Td>{item.children}</Table.Td>
        <Table.Td>{item.adult}</Table.Td>
        <Table.Td>{item.teen}</Table.Td>
        <Table.Td>{item.created_at}</Table.Td>
        <Table.Th>
          <Button size="xs" onClick={handleInsight}>Read More    <IconChevronRight size="0.9rem" stroke={1.5} /> </Button>
        </Table.Th>
      </Table.Tr>
    );
  });

  return (
    <ScrollArea>
      <Table miw={800} verticalSpacing="sm">
        <Table.Thead>
          <Table.Tr>
            <Table.Th style={{ width: rem(40) }}>
              <Checkbox
                onChange={toggleAll}
                checked={selection.length === data.length}
                indeterminate={selection.length > 0 && selection.length !== data.length}
              />
            </Table.Th>
            {/* <Table.Th>Author</Table.Th> */}
            <Table.Th>Author</Table.Th>
            <Table.Th>Service</Table.Th>
            <Table.Th>Children & Baby</Table.Th>
            <Table.Th>Adult</Table.Th>
            <Table.Th>Teens</Table.Th>
            <Table.Th>Date Created</Table.Th>
            <Table.Th>Option</Table.Th>
          </Table.Tr>
        </Table.Thead>
        <Table.Tbody>{rows}</Table.Tbody>
      </Table>
    </ScrollArea>
  );
}
