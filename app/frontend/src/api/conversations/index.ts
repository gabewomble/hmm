import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { ConversationService } from "../../../bindings/app/conversation_service";
import type { CreateConversationRequest } from "../../../bindings/app/conversation_service/models";

export function useConversations() {
	return useQuery({
		queryKey: ["conversations"],
		queryFn: () => ConversationService.ListConversations(),
	});
}

export function useCreateConversation() {
	const queryClient = useQueryClient();

	return useMutation({
		mutationFn: (input: CreateConversationRequest) =>
			ConversationService.CreateConversation(input),
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: ["conversations"] });
		},
	});
}
