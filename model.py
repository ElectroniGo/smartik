from datasets import load_dataset
from transformers import TrOCRProcessor, VisionEncoderDecoderModel, Seq2SeqTrainer, Seq2SeqTrainingArguments

# 1. Load IAM dataset (e.g. via CSV metadata)
dataset = load_dataset('csv', data_files={'train':'train.csv','validation':'val.csv','test':'test.csv'})

# 2. Load processor + pretrained model
processor = TrOCRProcessor.from_pretrained('microsoft/trocr-base-handwritten')
model = VisionEncoderDecoderModel.from_pretrained('microsoft/trocr-base-handwritten')

# 3. Preprocessing function
def preprocess(batch):
    images = [Image.open(path).convert("RGB") for path in batch['image_path']]
    pixel_values = processor(images=images, return_tensors="pt").pixel_values
    labels = processor.tokenizer(batch['text'], padding=True, truncation=True).input_ids
    batch['pixel_values'] = pixel_values
    batch['labels'] = labels
    return batch

tokenized = dataset.map(preprocess, batched=True)

# 4. Training setup
training_args = Seq2SeqTrainingArguments(
    output_dir='./trocr-iam-finetuned',
    per_device_train_batch_size=8,
    per_device_eval_batch_size=4,
    predict_with_generate=True,
    evaluation_strategy='steps',
    num_train_epochs=3,
    logging_steps=500,
    save_steps=1000,
)

trainer = Seq2SeqTrainer(
    model=model,
    args=training_args,
    train_dataset=tokenized['train'],
    eval_dataset=tokenized['validation'],
    tokenizer=processor.tokenizer,
)

trainer.train()
trainer.save_model('./trocr-iam-finetuned')
