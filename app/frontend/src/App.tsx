import { MantineProvider } from "@mantine/core";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

// core styles are required for all packages
import "@mantine/core/styles.css";
import AppRoutes from "./routes";

// other css files are required only if
// you are using components from the corresponding package
// import '@mantine/dates/styles.css';
// import '@mantine/dropzone/styles.css';
// import '@mantine/code-highlight/styles.css';

const queryClient = new QueryClient();

function App() {
	return (
		<QueryClientProvider client={queryClient}>
			<MantineProvider defaultColorScheme="auto">
				<AppRoutes />
			</MantineProvider>
		</QueryClientProvider>
	);
}

export default App;
