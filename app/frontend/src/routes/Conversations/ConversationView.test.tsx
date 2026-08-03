import { MantineProvider } from "@mantine/core";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import "@testing-library/jest-dom";
import { beforeEach, describe, expect, it, vi } from "vitest";
import ConversationView from "./ConversationView";

const mockNavigate = vi.fn();
let mockConversations: { id: string; name: string }[] = [
	{ id: "test-id", name: "Test Conversation" },
];

vi.mock("../../api/messages", () => ({
	useMessages: () => ({
		data: [],
		isLoading: false,
	}),
	useCreateMessage: () => ({
		mutate: vi.fn(),
		isPending: false,
	}),
}));

vi.mock("../../api/conversations", () => ({
	useConversations: () => ({
		data: mockConversations,
	}),
	useDeleteConversation: () => ({
		mutate: vi.fn(),
		isPending: false,
	}),
}));

vi.mock("react-router", () => ({
	useParams: () => ({ id: "test-id" }),
	useNavigate: () => mockNavigate,
}));

function renderWithProviders(ui: React.ReactElement) {
	const queryClient = new QueryClient();
	return render(
		<QueryClientProvider client={queryClient}>
			<MantineProvider>{ui}</MantineProvider>
		</QueryClientProvider>,
	);
}

describe("ConversationView", () => {
	beforeEach(() => {
		mockConversations = [{ id: "test-id", name: "Test Conversation" }];
	});

	it("renders conversation name in nav bar", () => {
		renderWithProviders(<ConversationView />);
		expect(screen.getByText("Test Conversation")).toBeInTheDocument();
	});

	it("renders fallback title when conversation not found", () => {
		mockConversations = [];

		renderWithProviders(<ConversationView />);
		expect(screen.getByText("Conversation")).toBeInTheDocument();
	});

	it("renders options menu button", () => {
		renderWithProviders(<ConversationView />);
		expect(
			screen.getByRole("button", { name: "Conversation options" }),
		).toBeInTheDocument();
	});

	it("opens menu and shows delete option when clicking options button", async () => {
		renderWithProviders(<ConversationView />);

		const optionsButton = screen.getByRole("button", {
			name: "Conversation options",
		});
		fireEvent.click(optionsButton);

		await waitFor(() => {
			expect(screen.getByText("Delete conversation")).toBeInTheDocument();
		});
	});

	it("renders empty state message", () => {
		renderWithProviders(<ConversationView />);
		expect(
			screen.getByText("No messages yet. Start the conversation!"),
		).toBeInTheDocument();
	});

	it("renders input field", () => {
		renderWithProviders(<ConversationView />);
		expect(
			screen.getByPlaceholderText("Type a message..."),
		).toBeInTheDocument();
	});
});
