import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { MessageService } from "../../../bindings/app/message_service";
import type { CreateMessageRequest } from "../../../bindings/app/message_service/models";

export function useMessages(conversationId: string | undefined) {
	return useQuery({
		queryKey: ["messages", conversationId],
		queryFn: () => {
			if (!conversationId) {
				throw new Error("conversationId is required");
			}
			return MessageService.ListMessagesByConversation({ conversationId });
		},
		enabled: !!conversationId,
	});
}

export function useCreateMessage() {
	const queryClient = useQueryClient();

	return useMutation({
		mutationFn: (input: CreateMessageRequest) =>
			MessageService.CreateMessage(input),
		onSuccess: (_data, variables) => {
			queryClient.invalidateQueries({
				queryKey: ["messages", variables.conversationId],
			});
		},
	});
}
