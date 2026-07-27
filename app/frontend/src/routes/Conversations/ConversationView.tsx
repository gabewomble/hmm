import { ActionIcon, Group, Stack, Text, Textarea } from "@mantine/core";
import { Send } from "lucide-react";
import { useEffect, useRef, useState } from "react";
import { useParams } from "react-router";
import { useCreateMessage, useMessages } from "../../api/messages";
import ChatMessage from "../../components/ChatMessage/ChatMessage";
import classes from "./ConversationView.module.css";

export default function ConversationView() {
	const { id } = useParams<{ id: string }>();
	const { data: messages, isLoading } = useMessages(id);
	const createMessage = useCreateMessage();
	const [input, setInput] = useState("");
	const viewportRef = useRef<HTMLDivElement>(null);

	// biome-ignore lint/correctness/useExhaustiveDependencies: Need to scroll when messages change
	useEffect(() => {
		if (viewportRef.current) {
			viewportRef.current.scrollTo({
				top: viewportRef.current.scrollHeight,
				behavior: "smooth",
			});
		}
	}, [messages?.length]);

	const handleSubmit = () => {
		if (!input.trim() || !id) return;

		createMessage.mutate(
			{ conversationId: id, body: input.trim() },
			{
				onSuccess: () => {
					setInput("");
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
			<div ref={viewportRef} className={classes.messagesArea}>
				{isLoading ? (
					<Text c="dimmed">Loading messages...</Text>
				) : !messages || messages.length === 0 ? (
					<Text c="dimmed">No messages yet. Start the conversation!</Text>
				) : (
					<Stack gap="md">
						{messages.map((msg) => (
							<ChatMessage
								key={msg.id}
								body={msg.body}
								messageType={msg.messageType}
								createdAt={msg.createdAt}
							/>
						))}
					</Stack>
				)}
			</div>

			<div className={classes.inputArea}>
				<Group gap="sm" align="flex-end">
					<Textarea
						className={classes.textarea}
						placeholder="Type a message..."
						value={input}
						onChange={(e) => setInput(e.currentTarget.value)}
						onKeyDown={handleKeyDown}
						autosize
						minRows={1}
						maxRows={5}
						disabled={createMessage.isPending}
					/>
					<ActionIcon
						size="lg"
						variant="filled"
						color="blue"
						onClick={handleSubmit}
						loading={createMessage.isPending}
						disabled={!input.trim()}
					>
						<Send size={20} />
					</ActionIcon>
				</Group>
			</div>
		</Stack>
	);
}
