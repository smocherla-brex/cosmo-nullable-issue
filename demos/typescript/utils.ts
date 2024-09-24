export const baseSchema = `
  directive @override(from: String!) on FIELD_DEFINITION

  directive @key(fields: _FieldSet!, resolvable: Boolean = true) repeatable on OBJECT
  
  directive @shareable on FIELD_DEFINITION | OBJECT
  
  scalar _FieldSet
`;