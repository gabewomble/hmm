import { MantineProvider } from "@mantine/core";
import { render, screen } from "@testing-library/react";
import "@testing-library/jest-dom";
import { describe, expect, it } from "vitest";
import ContentNavBar from "./ContentNavBar";

function renderWithMantine(ui: React.ReactElement) {
	return render(<MantineProvider>{ui}</MantineProvider>);
}

describe("ContentNavBar", () => {
	it("renders title", () => {
		renderWithMantine(<ContentNavBar title="Test Title" />);
		expect(screen.getByText("Test Title")).toBeInTheDocument();
	});

	it("renders title as ReactNode", () => {
		renderWithMantine(
			<ContentNavBar title={<span data-testid="custom-title">Custom</span>} />,
		);
		expect(screen.getByTestId("custom-title")).toBeInTheDocument();
	});

	it("renders actions when provided", () => {
		renderWithMantine(
			<ContentNavBar
				title="Title"
				actions={<button type="button">Action</button>}
			/>,
		);
		expect(screen.getByRole("button", { name: "Action" })).toBeInTheDocument();
	});

	it("does not render actions container when actions is not provided", () => {
		const { container } = renderWithMantine(<ContentNavBar title="Title" />);
		expect(container.querySelector("[class*='actions']")).toBeNull();
	});

	it("renders multiple actions", () => {
		renderWithMantine(
			<ContentNavBar
				title="Title"
				actions={
					<>
						<button type="button">Action 1</button>
						<button type="button">Action 2</button>
					</>
				}
			/>,
		);
		expect(
			screen.getByRole("button", { name: "Action 1" }),
		).toBeInTheDocument();
		expect(
			screen.getByRole("button", { name: "Action 2" }),
		).toBeInTheDocument();
	});
});
