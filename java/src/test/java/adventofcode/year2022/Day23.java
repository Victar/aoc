package adventofcode.year2022;

import adventofcode.BaseTest;
import lombok.Data;
import org.junit.Ignore;
import org.junit.Test;

import java.util.*;

public class Day23 extends BaseTest {

	public static final int DAY = 23;

	public static int SIZE_X_START = 0;
	public static int SIZE_Y_START = 0;
	public static int SIZE_X_END = Integer.MIN_VALUE;
	public static int SIZE_Y_END = Integer.MIN_VALUE;
	public static HashMap<String, Point> map = new HashMap<>();
	public static List<Point> list = new ArrayList<>();
	public static Point direction = Point.NORTH;
	boolean canMove = true;

	@Ignore @Test public void runDownloadInput() throws Exception {
		downloadInput(DAY);
	}

	@Test public void runBoth() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		for (int y = 0; y < data.size(); y++) {
			for (int x = 0; x < data.get(0).length(); x++) {
				if (data.get(y).charAt(x) == '#') {
					Point current = new Point(x, y, data.get(y).charAt(x) + "");
					map.put(current.getId(), current);
					list.add(current);
				}
			}
		}
		updateMinMax();
		int round = 0;
		while (canMove) {
			round = round + 1;
			list = doRound();
			direction = Point.getNextDirection(direction);
			updateMinMax();
			if (round == 10) {
				int answerSilver = (SIZE_Y_END - SIZE_Y_START + 1) * (SIZE_X_END - SIZE_X_START + 1) - list.size();
				System.out.println(answerSilver);
			}
		}
		System.out.println(round);
		//		drawState();
	}

	public void updateMinMax() {
		for (Point point : list) {
			SIZE_X_END = Math.max(SIZE_X_END, point.x);
			SIZE_Y_END = Math.max(SIZE_Y_END, point.y);
			SIZE_X_START = Math.min(SIZE_X_START, point.x);
			SIZE_Y_START = Math.min(SIZE_Y_START, point.y);
		}
	}

	public List<Point> doRound() {
		Map<Point, Point> oldNewMap = new HashMap<>();
		Map<Point, Integer> newCount = new HashMap<>();
		boolean canMoveInround = false;
		for (Point point : list) {
			boolean canMoveSingle = getNewPosition(point, oldNewMap, newCount);
			canMoveInround = canMoveInround || canMoveSingle;
		}
		canMove = canMoveInround;
		List<Point> listResult = new ArrayList<>();
		for (Map.Entry<Point, Point> entry : oldNewMap.entrySet()) {
			Point oldP = entry.getKey();
			Point newP = entry.getValue();
			int newPCount = newCount.get(newP);
			if (newPCount > 1) {
				listResult.add(oldP);
			} else {
				listResult.add(newP);
			}
		}
		return listResult;
	}

	public boolean getNewPosition(Point point, Map<Point, Point> oldNewMap, Map<Point, Integer> newCount) {
		boolean hasNeighbors = false;
		for (int x = -1; x <= 1; x++) {
			for (int y = -1; y <= 1; y++) {
				if (!(x == 0 && y == 0)) {
					Point neighboar = point.add(Point.of(x, y));
					if (list.contains(neighboar)) {
						hasNeighbors = true;
						break;
					}
				}
			}
		}
		if (!hasNeighbors) {
			oldNewMap.put(point, point);
			newCount.put(point, newCount.getOrDefault(point, 0) + 1);
			return false;
		}
		Point currentDirection = direction;
		for (int i = 0; i < Point.DIRECTIONS.size(); i++) {
			List<Point> toCheckList = currentDirection.directionToCheck();
			boolean canMove = true;
			for (Point toCheck : toCheckList) {
				if (list.contains(point.add(toCheck))) {
					canMove = false;
				}
			}
			if (canMove) {
				Point newPoint = point.add(currentDirection);
				oldNewMap.put(point, newPoint);
				newCount.put(newPoint, newCount.getOrDefault(newPoint, 0) + 1);
				return true;
			}
			currentDirection = Point.getNextDirection(currentDirection);
		}

		oldNewMap.put(point, point);
		newCount.put(point, newCount.getOrDefault(point, 0) + 1);
		return true;

	}

	public void drawState() {
		System.out.println("~~~~~");
		for (int y = SIZE_Y_START; y <= SIZE_Y_END + 1; y++) {
			StringBuilder sb = new StringBuilder();
			for (int x = SIZE_X_START; x <= SIZE_X_END + 1; x++) {
				if (list.contains(Point.of(x, y))) {
					sb.append("#");
				} else {
					sb.append(".");
				}

			}
			System.out.println(sb);
		}
	}

	@Data static class Point {

		public static Point NORTH = Point.of(0, -1, "NORTH");
		public static Point EAST = Point.of(1, 0, "EAST");
		public static Point WEST = Point.of(-1, 0, "WEST");
		public static Point SOUTH = Point.of(0, 1, "SOUTH");
		public static List<Point> DIRECTIONS = Arrays.asList(NORTH, SOUTH, WEST, EAST);

		int x;
		int y;
		String type; // "#" - elve, ".", " ";
		//		Point newPosition;

		public Point(int x, int y, String type) {
			this.x = x;
			this.y = y;
			this.type = type;
		}

		public static Point of(int x, int y) {
			return new Point(x, y, "");
		}

		public static Point of(int x, int y, String type) {
			return new Point(x, y, type);
		}

		public static Point getNextDirection(Point direction) {
			int index = DIRECTIONS.indexOf(direction);
			return DIRECTIONS.get((index + 1) % 4);
		}

		public List<Point> directionToCheck() {
			List<Point> directions = new ArrayList<>();
			boolean isNorth = this.equals(NORTH);
			boolean isEast = this.equals(EAST);
			boolean isWest = this.equals(WEST);
			boolean isSouth = this.equals(SOUTH);
			if (isSouth) {
				return List.of(Point.of(-1, 1), Point.of(0, 1), Point.of(1, 1));
			}
			if (isEast) {
				return List.of(Point.of(1, 1), Point.of(1, 0), Point.of(1, -1));
			}
			if (isNorth) {
				return List.of(Point.of(-1, -1), Point.of(0, -1), Point.of(1, -1));
			}
			if (isWest) {
				return List.of(Point.of(-1, -1), Point.of(-1, 0), Point.of(-1, 1));
			}
			return directions;
		}

		public Point add(Point p) {
			return Point.of(this.x + p.x, this.y + p.y);
		}

		public String getId() {
			return x + ":" + y;
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
			return "{" + "x=" + x + ", y=" + y + "}";
		}
	}

}
