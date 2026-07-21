import { MantineProvider, Button, Stack, Title, Text, Container } from "@mantine/core";
import { CreateProvider, ListProviders } from "../bindings/app/providerservice";

function App() {
  const handleListProviders = async () => {
    try {
      const result = await ListProviders();
      console.log("Providers:", result);
      alert(JSON.stringify(result, null, 2));
    } catch (err) {
      console.error(err);
      alert("Error: " + err);
    }
  };

  const handleCreateProvider = async () => {
    try {
      const result = await CreateProvider({
        name: "My Claude API",
        providerType: "anthropic",
        baseUrl: null,
        apiKey: "test-key",
      });
      console.log("Created:", result);
      alert("Created: " + JSON.stringify(result, null, 2));
    } catch (err) {
      console.error(err);
      alert("Error: " + err);
    }
  };

  return (
    <MantineProvider defaultColorScheme="auto">
      <Container size="sm" py="xl">
        <Stack gap="md">
          <Title order={2}>Provider Service Test</Title>
          <Text size="sm" c="dimmed">
            Open browser console to see results
          </Text>
          <Button onClick={handleCreateProvider}>Create Test Provider</Button>
          <Button onClick={handleListProviders} variant="light">
            List Providers
          </Button>
        </Stack>
      </Container>
    </MantineProvider>
  );
}

export default App;
