import { MantineProvider } from "@mantine/core";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { fireEvent, render, screen } from "@testing-library/react";
import "@testing-library/jest-dom";
import { describe, expect, it, vi } from "vitest";
import NewConversation from "./NewConversation";

vi.mock("../../api/conversations", () => ({
	useCreateConversation: () => ({
		mutate: vi.fn(),
		isPending: false,
	}),
}));

vi.mock("react-router", () => ({
	useNavigate: () => vi.fn(),
}));

function renderWithProviders(ui: React.ReactElement) {
	const queryClient = new QueryClient();
	return render(
		<QueryClientProvider client={queryClient}>
			<MantineProvider>{ui}</MantineProvider>
		</QueryClientProvider>,
	);
}

describe("NewConversation", () => {
	it("renders input field", () => {
		renderWithProviders(<NewConversation />);

		expect(
			screen.getByPlaceholderText("Start a new conversation..."),
		).toBeInTheDocument();
	});

	it("disables send button when input is empty", () => {
		renderWithProviders(<NewConversation />);

		const sendButton = screen.getByRole("button");
		expect(sendButton).toBeDisabled();
	});

	it("enables send button when input has text", () => {
		renderWithProviders(<NewConversation />);

		const input = screen.getByPlaceholderText("Start a new conversation...");
		fireEvent.change(input, { target: { value: "Hello" } });

		const sendButton = screen.getByRole("button");
		expect(sendButton).not.toBeDisabled();
	});
});
