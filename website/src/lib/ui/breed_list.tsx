import { Button, Card, Grid, TextInput, Group, Modal, Text } from "@mantine/core";
import { type BreedResponse } from "../dtos/responses";
import { useDisclosure } from "@mantine/hooks";
import { useState } from "react";
import type { AddBreed } from "../dtos/requests/breed";

interface BreedListProps {
    categoryId: number;
    breeds: Array<BreedResponse>;
    setBreedName: (id: number, name: string) => void;
    addBreed: (request: AddBreed) => Promise<void> | void;
    deleteBreed: (id: number) => void;
    isLoading: boolean;
}

export default function BreedList({ categoryId, breeds, setBreedName, addBreed, deleteBreed, isLoading }: BreedListProps) {
    // Modal for adding breeds.
    const [opened, { open, close }] = useDisclosure(false);
    const [modalTextState, setModalTextState] = useState("");

    function BreedCard({ value, onChange }: { value: string; onChange: (v: string) => void }) {
        return (
            <Card shadow="sm" padding="lg">
                <Card.Section>
                    <TextInput disabled={isLoading} value={value} onChange={(event) => onChange(event.currentTarget.value)} />
                </Card.Section>
            </Card>
        );
    }

    return (
        <>
            <Modal opened={opened} onClose={close} centered size="md" withCloseButton>
                <Card>
                    <Text>
                        New breed name.
                    </Text>
                    <TextInput value={modalTextState} onChange={event => setModalTextState(event.currentTarget.value)} />
                    <Button onClick={() => addBreed({name:modalTextState, category_id: categoryId})}>
                        Submit
                    </Button>
                </Card>
            </Modal>
            {breeds.map((breed) => (
                <Grid key={breed.id} mt="md" align="center">
                    <Grid.Col span={8}>
                        <BreedCard
                            value={breed.name}
                            onChange={(v) => setBreedName(breed.id, v)}
                        />
                    </Grid.Col>
                    <Grid.Col span={4}>
                        <Group>
                            <Button color="red" size="xs" onClick={() => deleteBreed(breed.id)}>
                                Remove
                            </Button>
                        </Group>
                    </Grid.Col>
                </Grid>
            ))}

            <Group mt="md">
                <Button onClick={() => open()} variant="filled" radius="lg" size="xs" disabled={isLoading}>Add Breed</Button>
            </Group>
        </>
    );
}