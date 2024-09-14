from enum import Enum
from uuid import uuid4, UUID

from pydantic import BaseModel


class Role(str, Enum):
    user = "user"
    admin = "admin"


class User(BaseModel):
    id: UUID | None = uuid4()
    email: str
    roles: set[Role] = [Role.user]
