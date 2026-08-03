import { Paper, Text } from "@mantine/core";
import { MessageType } from "#bindings/app/storage/models";
import classes from "./ChatMessage.module.css";

interface ChatMessageProps {
	body: string;
	messageType: MessageType;
	createdAt: string;
}

export default function ChatMessage({ body, messageType }: ChatMessageProps) {
	const isUser = messageType === MessageType.MessageTypeUser;

	return (
		<Paper
			className={`${classes.message} ${isUser ? classes.user : classes.llm}`}
			p="sm"
		>
			<Text>{body}</Text>
		</Paper>
	);
}
