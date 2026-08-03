import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import type { DeleteConversationRequest } from "#bindings/app/conversation_service";
import { MessageService } from "#bindings/app/message_service";
import type {
	CreateMessageRequest,
	MessageResponse,
} from "#bindings/app/message_service/models";
import type { ValueOf } from "#src/types/helpers";
import type { IMutationOptions } from "../types";

export const MESSAGE_QUERY_KEY = {
	LIST_MESSAGES_BY_CONVERSATION: "ListMessagesByConversation",
} as const;

export type MessageQueryKey = ValueOf<typeof MESSAGE_QUERY_KEY>;

export function useMessages(conversationId: string | undefined) {
	return useQuery({
		queryKey: [MESSAGE_QUERY_KEY.LIST_MESSAGES_BY_CONVERSATION, conversationId],
		queryFn: () => {
			if (!conversationId) {
				throw new Error("conversationId is required");
			}
			return MessageService.ListMessagesByConversation({ conversationId });
		},
		enabled: !!conversationId,
	});
}

export function useCreateMessage(
	options?: IMutationOptions<MessageResponse, DeleteConversationRequest>,
) {
	const invalidateConversationMessages = useInvalidateConversationMessages();

	return useMutation({
		mutationFn: (input: CreateMessageRequest) =>
			MessageService.CreateMessage(input),
		onSuccess: (data, variables) => {
			invalidateConversationMessages(variables.conversationId);
			options?.onSuccess?.(data, variables);
		},
		onError: options?.onError,
	});
}

function useInvalidateConversationMessages() {
	const queryClient = useQueryClient();

	return (conversationId: string) => {
		queryClient.invalidateQueries({
			queryKey: [
				MESSAGE_QUERY_KEY.LIST_MESSAGES_BY_CONVERSATION,
				conversationId,
			],
		});
	};
}
