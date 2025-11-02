import { Button, TextInput, Group, Modal, Text } from "@mantine/core";
import { BreedResponseZod, ErrorResponseZod, type BreedResponse, type ErrorResponse } from "../dtos/responses";
import { useDisclosure } from "@mantine/hooks";
import { useState } from "react";
import { type AddBreed } from "../dtos/requests/breed";
import { notifications } from "@mantine/notifications";
import { API_BASE_URL } from "../../App";

interface BreedListProps {
    categoryId: number;
    breedList: BreedResponse[];
    changeBreedList: (newVal: BreedResponse[]) => void,
    submitting: boolean;
    setSubmitting: (loading: boolean) => void,
}

function validateBreedName(breedName: string): string | null {
    return /^[a-z]+$/.test(breedName) && breedName.length < 20 ? null : "This is not a valid name";
}

export default function BreedList({ categoryId, breedList, changeBreedList, submitting, setSubmitting }: BreedListProps) {
    // Modal for adding breeds.
    const [opened, { open, close }] = useDisclosure(false);
    const [modalTextState, setModalTextState] = useState("");

    const setError = (message: string) => {
        notifications.show({
            title: "Something went wrong",
            message: message
        })
    }

    const setMessage = (message: string) => {
        notifications.show({
            title: "Success!",
            message: message
        })
    }

    // Add a new breed by POSTing to the API and appending the returned breed
    const addBreedRequest = async (request: AddBreed) => {
        setSubmitting(true);
        try {
            const res = await fetch(`${API_BASE_URL}/breed`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(request),
            });
            if (!res.ok) {
                const errorMessage = ErrorResponseZod.safeParse(await res.json())
                if (errorMessage.success) {
                    setError((errorMessage.data as ErrorResponse).message);
                }
                else setError("Unexpected result");
                return;
            }
            const breedResponseJson = await res.json();
            const breedResponse = BreedResponseZod.parse(breedResponseJson) as BreedResponse;
            changeBreedList([...breedList, breedResponse])
            setMessage('Breed added');
        } catch (err) {
            console.error('addBreed error', err);
            const message = err instanceof Error ? err.message : String(err);
            setError(message);
        } finally {
            setSubmitting(false);
        }
    };

    const removeBreedRequest = async (id: number) => {
        setSubmitting(true);
        try {
            const res = await fetch(`${API_BASE_URL}/breed/${id}`, {
                method: 'DELETE',
            });
            if (!res.ok) {
                const errorMessage = ErrorResponseZod.safeParse(await res.json())
                if (errorMessage.success) {
                    setError((errorMessage.data as ErrorResponse).message)
                }
                else setError("Unexpected result");
                return;
            }
            changeBreedList(breedList.filter((b) => b.id !== id))
            setMessage("Breed deleted.")
        } catch (err) {
            console.error('addBreed error', err);
            const message = err instanceof Error ? err.message : String(err);
            setError(message);
        } finally {
            setSubmitting(false);
        }

    };

    function BreedCard({ originalValue }: { originalValue: BreedResponse }) {
        const [breedNameEdited, setBreedText] = useState(originalValue.name)

        const updateBreedName = async (id: number, name: string) => {
            setSubmitting(true);
            try {
                const res = await fetch(`${API_BASE_URL}/breed/${id}`, {
                    method: 'PATCH',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify(
                        {
                            name: name,
                        }
                    ),
                });
                if (!res.ok) {
                    const errorMessage = ErrorResponseZod.safeParse(await res.json())
                    if (errorMessage.success) {
                        setError((errorMessage.data as ErrorResponse).message)
                        return;
                    }
                    else {
                        setError("Updating the breed name resulted in a failed value.");
                        return
                    }
                }
                const breedResponseJson = await res.json();
                const breedResponse = BreedResponseZod.parse(breedResponseJson) as BreedResponse;
                changeBreedList(breedList.map((breed) => breed.id !== id ? breed : { ...breed, name: breedResponse.name }))
                setMessage('Breed changed');
            } catch (err) {
                console.error('addBreed error', err);
                const message = err instanceof Error ? err.message : String(err);
                setError(message);
            } finally {
                setSubmitting(false);
            }
        };

        const [error, setError] = useState<string | null>(null)

        const handleInput = (event: React.ChangeEvent<HTMLInputElement>) => {
            setBreedText(event.target.value)
            setError(validateBreedName(event.target.value))
        }

        return (
            <>
                <TextInput disabled={submitting} value={breedNameEdited} onChange={handleInput} error={error} />
                {
                    breedNameEdited !== originalValue.name &&
                    <>
                        <Button disabled={error!==null} onClick={() => updateBreedName(originalValue.id, breedNameEdited)}>
                            Update
                        </Button>
                        <Button onClick={() => setBreedText(originalValue.name)}>
                            Reset
                        </Button>
                    </>
                }
            </>
        );
    }

    return (
        <>
            <Modal opened={opened} onClose={close} centered size="md" withCloseButton title={"Add new breed"}>
                <Group>
                    <Text>
                        New breed name:
                    </Text>
                    <TextInput value={modalTextState} onChange={event => setModalTextState(event.currentTarget.value)} />
                    <Button onClick={() => addBreedRequest({ name: modalTextState, category_id: categoryId })}>
                        Submit
                    </Button>
                </Group>
            </Modal>
            {breedList.map((breed) => (
                <Group>
                    <BreedCard
                        originalValue={breed}
                    />
                    <Button color="red" size="xs" onClick={() => removeBreedRequest(breed.id)}>
                        Remove
                    </Button>
                </Group>
            )
            )}
            <Button onClick={() => open()} variant="filled" radius="lg" size="xs" disabled={submitting}>Add Breed</Button>
        </>
    );
}