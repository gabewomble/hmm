import { MantineProvider } from "@mantine/core";
import { render, screen } from "@testing-library/react";
import "@testing-library/jest-dom";
import { describe, expect, it } from "vitest";
import { MessageType } from "#bindings/app/storage/models";
import ChatMessage from "./ChatMessage";

function renderWithMantine(ui: React.ReactElement) {
	return render(<MantineProvider>{ui}</MantineProvider>);
}

describe("ChatMessage", () => {
	it("renders user message", () => {
		renderWithMantine(
			<ChatMessage
				body="Hello world"
				messageType={MessageType.MessageTypeUser}
				createdAt="2024-01-01T00:00:00Z"
			/>,
		);

		expect(screen.getByText("Hello world")).toBeInTheDocument();
	});

	it("renders llm message", () => {
		renderWithMantine(
			<ChatMessage
				body="I am an AI response"
				messageType={MessageType.MessageTypeLLM}
				createdAt="2024-01-01T00:00:00Z"
			/>,
		);

		expect(screen.getByText("I am an AI response")).toBeInTheDocument();
	});
});
