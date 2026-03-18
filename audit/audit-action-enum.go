package audit

type Action string

const (
    ActionCreate Action = "create"
    ActionDelete Action = "delete"
    ActionUpload Action = "upload"
    ActionSave   Action = "save"
    ActionMove   Action = "move"
    ActionCopy   Action = "copy"
    ActionRename Action = "rename"
)
