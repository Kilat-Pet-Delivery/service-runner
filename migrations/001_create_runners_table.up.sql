CREATE TABLE runners (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID UNIQUE NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    vehicle_type VARCHAR(20) NOT NULL,
    vehicle_plate VARCHAR(20),
    vehicle_model VARCHAR(100),
    vehicle_year INT,
    air_conditioned BOOLEAN DEFAULT FALSE,
    session_status VARCHAR(20) NOT NULL DEFAULT 'inactive',
    current_lat DOUBLE PRECISION,
    current_lng DOUBLE PRECISION,
    rating DECIMAL(3,2) DEFAULT 0.0,
    total_trips INT DEFAULT 0,
    completion_rate DECIMAL(3,2) DEFAULT 1.0,
    version BIGINT NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_runners_user ON runners(user_id);
CREATE INDEX idx_runners_status ON runners(session_status);
