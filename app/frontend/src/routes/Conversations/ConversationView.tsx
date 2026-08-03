import { ActionIcon, Group, Menu, Stack, Text, Textarea } from "@mantine/core";
import { MoreVertical, Send, Trash2 } from "lucide-react";
import { useState } from "react";
import { useNavigate, useParams } from "react-router";
import {
	useConversations,
	useDeleteConversation,
} from "../../api/conversations";
import { useCreateMessage, useMessages } from "../../api/messages";
import ChatMessage from "../../components/ChatMessage/ChatMessage";
import ContentNavBar from "../../components/ContentNavBar";
import classes from "./ConversationView.module.css";

export default function ConversationView() {
	const navigate = useNavigate();

	const { id } = useParams<{ id: string }>();

	const { data: messages, isLoading } = useMessages(id);
	const { data: conversations } = useConversations();

	const createMessage = useCreateMessage();
	const deleteConversation = useDeleteConversation();

	const [input, setInput] = useState("");

	const conversation = conversations?.find((c) => c.id === id);

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

	const handleDelete = () => {
		if (!id) return;
		deleteConversation.mutate(
			{ conversationId: id },
			{
				onSuccess: () => {
					navigate("/conversations/new");
				},
			},
		);
	};

	const actions = (
		<Menu position="bottom-end" shadow="md">
			<Menu.Target>
				<ActionIcon variant="subtle" aria-label="Conversation options">
					<MoreVertical size={20} />
				</ActionIcon>
			</Menu.Target>
			<Menu.Dropdown>
				<Menu.Item
					color="red"
					leftSection={<Trash2 size={16} />}
					onClick={handleDelete}
				>
					Delete conversation
				</Menu.Item>
			</Menu.Dropdown>
		</Menu>
	);

	return (
		<Stack className={classes.container} gap={0}>
			<ContentNavBar
				title={conversation?.name ?? "Conversation"}
				actions={actions}
			/>

			<div className={classes.messagesArea}>
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
