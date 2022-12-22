package adventofcode.year2022;

import adventofcode.BaseTest;
import lombok.Data;
import org.apache.commons.lang3.StringUtils;
import org.junit.Ignore;
import org.junit.Test;

import java.util.*;

public class Day22 extends BaseTest {

	public static final int DAY = 22;

	public static int SIZE_X = Integer.MIN_VALUE;
	public static int SIZE_Y = Integer.MIN_VALUE;
	public static int SIZE_CUBE = Integer.MIN_VALUE;

	Map<String, Point> map;
	Point position = Point.of(0, 0, "x");
	Point direction = Point.RIGHT;

	@Ignore @Test public void runDownloadInput() throws Exception {
		downloadInput(DAY);
	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		String instruction = data.get(data.size() - 1);
		for (int i = 0; i < data.size() - 1; i++) {
			SIZE_X = Math.max(SIZE_X, data.get(i).length());
		}
		SIZE_Y = data.size() - 2;
		System.out.println("SIZE_X:" + SIZE_X + " SIZE_Y:" + SIZE_Y);
		System.out.println(instruction);
		map = new HashMap<>();
		for (int y = 0; y < data.size() - 2; y++) {
			for (int x = 0; x < SIZE_X; x++) {
				String name = " ";
				if (data.get(y).length() > x) {
					name = data.get(y).charAt(x) + "";
				}
				Point current = new Point(x, y, name);
				map.put(current.getId(), current);
			}
		}
		Point direction = Point.RIGHT;
		String current = StringUtils.EMPTY;
		int i = 0;
		for (; i < instruction.length(); i++) {
			char cur = instruction.charAt(i);
			if (cur == 'R' || cur == 'L') {
				doInstruction(direction, Integer.parseInt(current));
				direction = Point.getNextInstruction(direction, cur);
				current = StringUtils.EMPTY;
			} else {
				current = current + cur;
			}
		}
		doInstruction(direction, Integer.parseInt(current));
		int bonusDirection = 0;
		if (Point.DOWN == direction) {
			bonusDirection = 1;
		}
		if (Point.LEFT == direction) {
			bonusDirection = 2;
		}
		if (Point.UP == direction) {
			bonusDirection = 3;
		}
		System.out.println(position + "  " + direction);
		System.out.println(4 * (position.x + 1) + 1000 * (position.y + 1) + bonusDirection);
		drawState();
	}

	public void doInstructionGold(int steps) {
		int cur = 0;
		Point lastPosition = position;
		Point lastDirection = direction;

		for (; cur < steps; ) {
			PositionDirection next = position.nextPointOnCube(direction);
			Point nextOnMap = map.get(next.position.getId());
			if (nextOnMap.isWall()) {
				position = lastPosition;
				direction = lastDirection;
				break;
			}
			if (nextOnMap.isSpace()) {
				throw new IllegalStateException("This should not happen on gold part");
			}
			if (nextOnMap.isField()) {
				cur++;
				position = next.position;
				direction = next.direction;
				lastPosition = position;
				lastDirection = direction;
			}
		}
		//		drawState();
	}

	public void doInstruction(Point instruction, int steps) {
		int cur = 0;
		Point lastField = position;
		for (; cur < steps; ) {
			Point next = position.nextPoint(instruction);
			Point nextOnMap = map.get(next.getId());
			if (nextOnMap.isWall()) {
				position = lastField;
				break;
			}
			if (nextOnMap.isSpace()) {
				position = next;
				continue;
			}
			if (nextOnMap.isField()) {
				cur++;
				position = next;
				lastField = position;

			}
		}
		//		drawState();
	}

	public void drawState() {
		System.out.println("      ");
		System.out.println(position);
		System.out.println("      ");

		for (int y = 0; y < SIZE_Y; y++) {
			StringBuilder sb = new StringBuilder();
			for (int x = 0; x < SIZE_X; x++) {
				String key = x + ":" + y;
				if (key.equals(position.getId())) {
					sb.append("x");
				} else {
					Point cube = map.get(key).getCube();
					if (cube != null) {
						sb.append(cube.type);
					} else {
						sb.append(map.get(key).type);
					}
				}
			}
			System.out.println(sb);
		}
	}

	public void doInstructionSilver(Point instruction, int steps) {
		//		System.out.println(steps + "  " + instruction.type  + " " + position);
		//		drawState();
		int cur = 0;
		Point lastField = position;
		for (; cur < steps; ) {
			Point next = position.nextPoint(instruction);
			Point nextOnMap = map.get(next.getId());
			if (nextOnMap.isWall()) {
				position = lastField;
				break;
			}
			if (nextOnMap.isSpace()) {
				position = next;
				continue;
			}
			if (nextOnMap.isField()) {
				cur++;
				position = next;
				lastField = position;

			}
		}
		//		drawState();
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		String instruction = data.get(data.size() - 1);
		for (int i = 0; i < data.size() - 1; i++) {
			SIZE_X = Math.max(SIZE_X, data.get(i).length());
		}
		SIZE_Y = data.size() - 2;
		SIZE_CUBE = SIZE_Y / 4;
		position = Point.of(SIZE_CUBE, 0, "x"); // override for test
		System.out.println("SIZE_X:" + SIZE_X + " SIZE_Y:" + SIZE_Y + " SIZE_CUBE:" + SIZE_CUBE);
		System.out.println(instruction);
		map = new HashMap<>();
		for (int y = 0; y < data.size() - 2; y++) {
			for (int x = 0; x < SIZE_X; x++) {
				String name = " ";
				if (data.get(y).length() > x) {
					name = data.get(y).charAt(x) + "";
				}
				//For test
				if (name.equals("x")) {
					position = Point.of(x, y, "x");
					name = ".";
				}
				//For test end
				Point current = new Point(x, y, name);
				map.put(current.getId(), current);
			}
		}
		drawState();
		direction = Point.RIGHT;
		String testDirection = data.get(data.size() - 2);
		if (StringUtils.isNotEmpty(testDirection)) {
			direction = Point.DIRECTIONS.get(Integer.parseInt(testDirection));
		}

		String current = StringUtils.EMPTY;
		int i = 0;
		for (; i < instruction.length(); i++) {
			char cur = instruction.charAt(i);
			if (cur == 'R' || cur == 'L') {
				doInstructionGold(Integer.parseInt(current));
				direction = Point.getNextInstruction(direction, cur);
				current = StringUtils.EMPTY;
			} else {
				current = current + cur;
			}
		}
		doInstructionGold(Integer.parseInt(current));
		int bonus = Integer.parseInt(position.getCube().type);
		System.out.println(position + "  " + bonus);
		System.out.println(4 * (position.x + 1) + 1000 * (position.y + 1) + bonus);
		drawState();
	}

	@Data static class PositionDirection {

		Point position;
		Point direction;

		public static PositionDirection of(Point position, Point direction) {
			PositionDirection pd = new PositionDirection();
			pd.position = position;
			pd.direction = direction;
			return pd;
		}
	}

	@Data static class Point {

		public static Point RIGHT = Point.of(1, 0, "RIGHT");
		public static Point LEFT = Point.of(-1, 0, "LEFT");
		public static Point DOWN = Point.of(0, 1, "DOWN");
		public static Point UP = Point.of(0, -1, "UP");

		public static Point CUBE_O = Point.of(0, 0, "~");
		public static Point CUBE_1 = Point.of(1, 0, "1");
		public static Point CUBE_2 = Point.of(2, 0, "2");
		public static Point CUBE_3 = Point.of(1, 1, "3");
		public static Point CUBE_4 = Point.of(0, 2, "4");
		public static Point CUBE_5 = Point.of(1, 2, "5");
		public static Point CUBE_6 = Point.of(0, 3, "6");

		//		public static Point CUBE_1 = Point.of(2, 0, "1");
		//		public static Point CUBE_2 = Point.of(0, 1, "2");
		//		public static Point CUBE_3 = Point.of(1, 1, "3");
		//		public static Point CUBE_4 = Point.of(2, 1, "4");
		//		public static Point CUBE_5 = Point.of(2, 2, "5");
		//		public static Point CUBE_6 = Point.of(3, 2, "6");

		public static List<Point> CUBES = Arrays.asList(CUBE_1, CUBE_2, CUBE_3, CUBE_4, CUBE_5, CUBE_6);

		public static List<Point> DIRECTIONS = Arrays.asList(RIGHT, DOWN, LEFT, UP);
		int x;
		int y;
		String type; // "#", ".", " ";

		public Point(int x, int y, String type) {
			this.x = x;
			this.y = y;
			this.type = type;
		}

		public static Point getNextInstruction(Point current, char direction) {
			int index = DIRECTIONS.indexOf(current);
			if ('R' == direction) {
				index = index + 1;
			}
			if ('L' == direction) {
				index = index - 1;
			}
			return DIRECTIONS.get((index + 4) % 4);
		}

		public static Point of(int x, int y) {
			return new Point(x, y, "");
		}

		public static Point of(int x, int y, String type) {
			return new Point(x, y, type);
		}

		public boolean isSpace() {
			return " ".equals(this.type);
		}

		public boolean isWall() {
			return "#".equals(this.type);
		}

		public boolean isField() {
			return ".".equals(this.type);
		}

		public Point nextPoint(Point direction) {
			int xNew = (this.x + direction.x + SIZE_X) % SIZE_X;
			int yNew = (this.y + direction.y + SIZE_Y) % SIZE_Y;
			return Point.of(xNew, yNew, this.type);
		}

		public String getId() {
			return x + ":" + y;
		}

		public Point getCube() {
			Point currentCube = Point.of(this.x / SIZE_CUBE, this.y / SIZE_CUBE);
			for (Point cube : CUBES) {
				if (cube.equals(currentCube)) {
					return cube;
				}
			}
			return CUBE_O;
		}

		@Override public boolean equals(Object o) {
			if (this == o) return true;
			if (o == null || getClass() != o.getClass()) return false;
			Point point = (Point) o;
			return x == point.x && y == point.y;
		}

		@Override public int hashCode() {
			return Objects.hash(x, y);
		}

		@Override public String toString() {
			return "{" + "x=" + x + ", y=" + y + ", type='" + type + '\'' + '}';
		}

		private PositionDirection cubeGetNewPoint(int x, int y, Point direction) {
			Point position = new Point(this.x * SIZE_CUBE + x, this.y * SIZE_CUBE + y, "x");
			if (UP.equals(direction)) {
				return PositionDirection.of(position, DOWN);
			}
			if (DOWN.equals(direction)) {
				return PositionDirection.of(position, UP);
			}
			if (LEFT.equals(direction)) {
				return PositionDirection.of(position, RIGHT);
			}
			if (RIGHT.equals(direction)) {
				return PositionDirection.of(position, LEFT);
			}
			return null;
		}

		public PositionDirection nextPointOnCube(Point direction) {
			Point currentCube = getCube();
			Point nextPoint = nextPoint(direction);
			Point nextPointCube = nextPoint.getCube();
			if (nextPointCube.equals(currentCube)) { //all good in scope of one cube;
				return PositionDirection.of(nextPoint, direction);
			}
			//Cube was changed

			//Current cube;
			boolean isCube1 = CUBE_1.equals(currentCube);
			boolean isCube2 = CUBE_2.equals(currentCube);
			boolean isCube3 = CUBE_3.equals(currentCube);
			boolean isCube4 = CUBE_4.equals(currentCube);
			boolean isCube5 = CUBE_5.equals(currentCube);
			boolean isCube6 = CUBE_6.equals(currentCube);

			int xCube = this.x % SIZE_CUBE;
			int yCube = this.y % SIZE_CUBE;

			System.out.println(xCube + "  " + yCube);

			int SIZE = SIZE_CUBE - 1;
			if (isCube1) {
				if (direction.equals(UP)) {
					return CUBE_6.cubeGetNewPoint(0, xCube, LEFT); // check (/)
				}
				if (direction.equals(DOWN)) {
					return CUBE_3.cubeGetNewPoint(xCube, 0, UP);
				}
				if (direction.equals(LEFT)) {
					return CUBE_4.cubeGetNewPoint(0, SIZE - yCube, LEFT);
				}
				if (direction.equals(RIGHT)) {
					return CUBE_2.cubeGetNewPoint(0, yCube, LEFT);
				}
			}
			if (isCube2) {
				if (direction.equals(UP)) {
					return CUBE_6.cubeGetNewPoint(xCube, SIZE, DOWN); // check (/)
				}
				if (direction.equals(DOWN)) {
					return CUBE_3.cubeGetNewPoint(SIZE, xCube, RIGHT);
				}
				if (direction.equals(LEFT)) {
					return CUBE_1.cubeGetNewPoint(SIZE, yCube, RIGHT);
				}
				if (direction.equals(RIGHT)) {
					return CUBE_5.cubeGetNewPoint(SIZE, SIZE - yCube, RIGHT);
				}
			}
			//tested
			if (isCube3) {
				if (direction.equals(UP)) {
					return CUBE_1.cubeGetNewPoint(xCube, SIZE, DOWN);
				}
				if (direction.equals(DOWN)) {
					return CUBE_5.cubeGetNewPoint(xCube, 0, UP);
				}
				if (direction.equals(LEFT)) {
					return CUBE_4.cubeGetNewPoint(yCube, 0, UP);
				}
				if (direction.equals(RIGHT)) {
					return CUBE_2.cubeGetNewPoint(yCube, SIZE, DOWN);
				}
			}
			if (isCube4) {
				if (direction.equals(UP)) {
					return CUBE_3.cubeGetNewPoint(0, xCube, LEFT);
				}
				if (direction.equals(DOWN)) {
					return CUBE_6.cubeGetNewPoint(xCube, 0, UP);
				}
				if (direction.equals(LEFT)) {
					return CUBE_1.cubeGetNewPoint(0, SIZE - yCube, LEFT);
				}
				if (direction.equals(RIGHT)) {
					return CUBE_5.cubeGetNewPoint(0, yCube, LEFT);
				}
			}
			if (isCube5) {
				if (direction.equals(UP)) {
					return CUBE_3.cubeGetNewPoint(xCube, SIZE, DOWN);
				}
				if (direction.equals(DOWN)) {
					return CUBE_6.cubeGetNewPoint(SIZE, xCube, RIGHT);
				}
				if (direction.equals(LEFT)) {
					return CUBE_4.cubeGetNewPoint(SIZE, yCube, RIGHT);
				}
				if (direction.equals(RIGHT)) {
					return CUBE_2.cubeGetNewPoint(SIZE, SIZE - yCube, RIGHT);
				}
			}
			if (isCube6) {
				if (direction.equals(UP)) {
					return CUBE_4.cubeGetNewPoint(xCube, SIZE, DOWN);
				}
				if (direction.equals(DOWN)) {
					return CUBE_2.cubeGetNewPoint(xCube, 0, UP);
				}
				if (direction.equals(LEFT)) {
					return CUBE_1.cubeGetNewPoint(yCube, 0, UP);
				}
				if (direction.equals(RIGHT)) {
					return CUBE_5.cubeGetNewPoint(yCube, SIZE, DOWN);
				}
			}

			return null;
		}

		public PositionDirection nextPointOnCubeTestInput(Point direction) {
			Point currentCube = getCube();
			Point nextPoint = nextPoint(direction);
			Point nextPointCube = nextPoint.getCube();
			if (nextPointCube.equals(currentCube)) { //all good in scope of one cube;
				return PositionDirection.of(nextPoint, direction);
			}
			//Cube was changed

			//Current cube;
			boolean isCube1 = CUBE_1.equals(currentCube);
			boolean isCube2 = CUBE_2.equals(currentCube);
			boolean isCube3 = CUBE_3.equals(currentCube);
			boolean isCube4 = CUBE_4.equals(currentCube);
			boolean isCube5 = CUBE_5.equals(currentCube);
			boolean isCube6 = CUBE_6.equals(currentCube);

			int xCube = this.x % SIZE_CUBE;
			int yCube = this.y % SIZE_CUBE;

			System.out.println(xCube + "  " + yCube);

			int SIZE = SIZE_CUBE - 1;
			if (isCube1) {
				if (direction.equals(DOWN)) {
					return CUBE_4.cubeGetNewPoint(xCube, 0, UP);
				}
				if (direction.equals(UP)) {
					return CUBE_2.cubeGetNewPoint(SIZE - xCube, 0, UP); // check (/)
				}
				if (direction.equals(LEFT)) {
					return CUBE_3.cubeGetNewPoint(yCube, 0, UP);
				}
				if (direction.equals(RIGHT)) {
					return CUBE_6.cubeGetNewPoint(SIZE, SIZE - yCube, RIGHT);
				}
			}

			if (isCube2) {
				if (direction.equals(DOWN)) {
					return CUBE_5.cubeGetNewPoint(SIZE - xCube, SIZE, DOWN);
				}
				if (direction.equals(UP)) {
					return CUBE_1.cubeGetNewPoint(SIZE - xCube, 0, UP); // check (/)
				}
				if (direction.equals(LEFT)) {
					return CUBE_6.cubeGetNewPoint(SIZE - yCube, SIZE, DOWN);
				}
				if (direction.equals(RIGHT)) {
					return CUBE_3.cubeGetNewPoint(0, yCube, LEFT);
				}
			}
			if (isCube3) {
				if (direction.equals(DOWN)) {
					return CUBE_5.cubeGetNewPoint(0, SIZE - xCube, LEFT);
				}
				if (direction.equals(UP)) {
					return CUBE_1.cubeGetNewPoint(0, xCube, LEFT);
				}
				if (direction.equals(LEFT)) {
					return CUBE_2.cubeGetNewPoint(SIZE, yCube, RIGHT);
				}
				if (direction.equals(RIGHT)) {
					return CUBE_4.cubeGetNewPoint(0, yCube, LEFT);
				}
			}
			if (isCube4) {
				if (direction.equals(UP)) {
					return CUBE_1.cubeGetNewPoint(xCube, SIZE, DOWN);
				}
				if (direction.equals(DOWN)) {
					return CUBE_5.cubeGetNewPoint(xCube, 0, UP);
				}
				if (direction.equals(LEFT)) {
					return CUBE_3.cubeGetNewPoint(SIZE, yCube, RIGHT);
				}
				if (direction.equals(RIGHT)) {
					return CUBE_6.cubeGetNewPoint(SIZE - yCube, 0, UP);
				}
			}
			if (isCube5) {
				if (direction.equals(UP)) {
					return CUBE_4.cubeGetNewPoint(xCube, SIZE, DOWN);
				}
				if (direction.equals(DOWN)) {
					return CUBE_2.cubeGetNewPoint(SIZE - xCube, SIZE, DOWN);
				}
				if (direction.equals(LEFT)) {
					return CUBE_3.cubeGetNewPoint(SIZE - yCube, SIZE, DOWN);
				}
				if (direction.equals(RIGHT)) {
					return CUBE_6.cubeGetNewPoint(0, yCube, LEFT);
				}
			}
			if (isCube6) {
				if (direction.equals(UP)) {
					return CUBE_4.cubeGetNewPoint(SIZE, SIZE - xCube, RIGHT);
				}
				if (direction.equals(DOWN)) {
					return CUBE_2.cubeGetNewPoint(0, SIZE - xCube, LEFT);
				}
				if (direction.equals(LEFT)) {
					return CUBE_5.cubeGetNewPoint(SIZE, yCube, RIGHT);
				}
				if (direction.equals(RIGHT)) {
					return CUBE_1.cubeGetNewPoint(SIZE, SIZE - yCube, RIGHT);
				}
			}

			return null;
		}

	}

}


