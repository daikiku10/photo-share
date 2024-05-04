"use client";

import type { ReactNode } from "react";
import { ModalContainer } from "@/components/ModalContainer";
import { PhotoCreateForm } from "../PhotoCreateForm";
import { PhotoCreateModal } from "../PhotoCreateModal";

type Props = {
  children: ReactNode;
  toggleClassName?: string;
};

export function PhotoCreateModalContainer({
  children,
  toggleClassName,
}: Props) {
  return (
    <ModalContainer
      toggleClassName={toggleClassName}
      content={(closeModal) => (
        <PhotoCreateModal close={closeModal}>
          <PhotoCreateForm />
        </PhotoCreateModal>
      )}
    >
      {children}
    </ModalContainer>
  );
}
