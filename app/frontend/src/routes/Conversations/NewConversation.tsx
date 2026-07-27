import { ActionIcon, Group, Stack, Text, Textarea } from "@mantine/core";
import { Send } from "lucide-react";
import { useState } from "react";
import { useNavigate } from "react-router";
import { useCreateConversation } from "../../api/conversations";
import { useCreateMessage } from "../../api/messages";
import classes from "./NewConversation.module.css";

export default function NewConversation() {
	const navigate = useNavigate();
	const createConversation = useCreateConversation();
	const createMessage = useCreateMessage();
	const [input, setInput] = useState("");
	const [error, setError] = useState<string | null>(null);
	const isPending = createConversation.isPending || createMessage.isPending;

	const handleSubmit = () => {
		if (!input.trim()) return;

		setError(null);
		createConversation.mutate(
			{ name: "New Conversation" },
			{
				onSuccess: (conversation) => {
					createMessage.mutate(
						{ conversationId: conversation.id, body: input.trim() },
						{
							onSuccess: () => {
								navigate(`/conversations/${conversation.id}`);
							},
							onError: () => {
								setError("Failed to send message. Please try again.");
								navigate(`/conversations/${conversation.id}`);
							},
						},
					);
				},
				onError: () => {
					setError("Failed to create conversation. Please try again.");
				},
			},
		);
	};

	const handleKeyDown = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
		if (e.key === "Enter" && !e.shiftKey) {
			e.preventDefault();
			handleSubmit();
		}
	};

	return (
		<Stack className={classes.container} gap={0}>
			<div className={classes.spacer} />

			{error && (
				<div className={classes.errorArea}>
					<Text c="red" size="sm">
						{error}
					</Text>
				</div>
			)}

			<div className={classes.inputArea}>
				<Group gap="sm" align="flex-end">
					<Textarea
						className={classes.textarea}
						placeholder="Start a new conversation..."
						value={input}
						onChange={(e) => setInput(e.currentTarget.value)}
						onKeyDown={handleKeyDown}
						autosize
						minRows={1}
						maxRows={5}
						disabled={isPending}
					/>
					<ActionIcon
						size="lg"
						variant="filled"
						color="blue"
						onClick={handleSubmit}
						loading={isPending}
						disabled={!input.trim()}
					>
						<Send size={20} />
					</ActionIcon>
				</Group>
			</div>
		</Stack>
	);
}
