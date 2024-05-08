import { useEffect, useId, useState } from "react";
import { Button } from "@/components/Button";
import { Label } from "@/components/Label";
import { TextArea } from "@/components/TextArea";
import { TextField } from "@/components/TextField";
import styles from "./style.module.css";

type Props = {
  onChange: (state: State) => void;
};
type State = {
  title: string;
  description: string;
  categoryId: string;
};

export function PhotoMeta({ onChange }: Props) {
  const componentId = useId();
  const titleId = `${componentId}-title`;
  const descriptionId = `${componentId}-description`;

  const [state, setState] = useState<State>({
    title: "",
    description: "",
    categoryId: "",
  });

  useEffect(() => {
    onChange(state);
  }, [onChange, state]);

  return (
    <div className={styles.meta}>
      <div className={styles.row}>
        <Label size="xsmall" htmlFor={titleId}>
          タイトル
        </Label>
        <TextField
          className={styles.title}
          id={titleId}
          value={state.title}
          onChange={(event) => {
            setState({ ...state, title: event.target.value });
          }}
        />
      </div>
      <div className={styles.row}>
        <Label size="xsmall" htmlFor={descriptionId}>
          写真の概要
        </Label>
        <TextArea
          className={styles.description}
          id={descriptionId}
          onChange={(event) => {
            setState({ ...state, description: event.target.value });
          }}
        />
      </div>
      <Button type="submit" color="orange" size="large">
        写真を投稿する
      </Button>
    </div>
  );
}
